package jwter

import (
	_ "embed"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// 秘密鍵と公開鍵の生成
// openssl genrsa 4096 > secret.pem
// openssl rsa -pubout < secret.pem  > public.pem

//go:embed testcert/secret.pem
var rawPrivateKey []byte

//go:embed testcert/public.pem
var rawPublicKey []byte

func TestGenerateAndParseToken(t *testing.T) {
	jtr, err := NewJWTer("testissuer", rawPrivateKey, rawPublicKey)
	if err != nil {
		t.Fatalf("fail to create JWTer: %s", err)
	}

	claims := []Claim{
		{Key: "color", Value: "darkblue"},
		{Key: "size", Value: "small"},
	}
	jwtid, token, err := jtr.GenerateToken("testtoken", time.Minute*10, claims)
	if err != nil {
		t.Fatalf("fail to create token: %s", err)
	}
	if len(jwtid) != 36 {
		t.Errorf("jwtid=uuid length is not 32+4, got=%d", len(jwtid))
	}

	r := httptest.NewRequest(http.MethodPost, "/", nil)
	r.Header.Set("Authorization", "Bearer "+string(token))
	parsed, err := jtr.ParseRequest(r)
	if err != nil {
		t.Fatalf("fail to parse request %s", err)
	}

	if parsed.JwtID() != jwtid {
		t.Errorf("jwtid=%s, but parsed=%s", jwtid, parsed.JwtID())
	}

	parsedclaims := parsed.PrivateClaims()
	for _, c := range claims {
		v, ok := parsedclaims[c.Key]
		if !ok || v.(string) != c.Value {
			t.Errorf("claim %s is lost", c.Key)
		}
	}
}

func TestTokenExpired(t *testing.T) {
	jtr, err := NewJWTer("testissuer", rawPrivateKey, rawPublicKey)
	if err != nil {
		t.Fatalf("fail to create JWTer: %s", err)
	}

	_, token, err := jtr.GenerateToken("testtoken", time.Second*1, nil)
	if err != nil {
		t.Fatalf("fail to create token: %s", err)
	}

	time.Sleep(time.Second * 2)

	r := httptest.NewRequest(http.MethodPost, "/", nil)
	r.Header.Set("Authorization", "Bearer "+string(token))
	_, err = jtr.ParseRequest(r)

	if err == nil {
		t.Fatalf("fail to validate expiration")
	}

	if fmt.Sprintf("%s", err) != `failed to validate token: "exp" not satisfied` {
		t.Errorf("unmatch expiration message: %s", err)
	}
}
