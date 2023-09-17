package jwter

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"net/http"
	"time"
)

type JWTer struct {
	Issuer     string
	PrivateKey jwk.Key
	PublicKey  jwk.Key
}

type Claim struct {
	Key   string
	Value string
}

func NewJWTer(issuer string, rawPrivateKey, rawPublicKey []byte) (*JWTer, error) {
	jwter := &JWTer{Issuer: issuer}

	privkey, err := jwk.ParseKey(rawPrivateKey, jwk.WithPEM(true))
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key")
	}
	pubkey, err := jwk.ParseKey(rawPublicKey, jwk.WithPEM(true))
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key")
	}

	jwter.PrivateKey = privkey
	jwter.PublicKey = pubkey

	return jwter, nil
}

func (j *JWTer) GenerateToken(subject string, duration time.Duration, claims []Claim) (string, []byte, error) {

	now := time.Now().In(time.FixedZone("Asia/Tokyo", 9*60*60))
	builder := jwt.NewBuilder().
		JwtID(uuid.New().String()).
		Issuer(j.Issuer).
		Subject(subject).
		IssuedAt(now).
		Expiration(now.Add(duration))

	for _, c := range claims {
		builder = builder.Claim(c.Key, c.Value)
	}

	token, err := builder.Build()

	if err != nil {
		return "", nil, fmt.Errorf("failed to build token: %w", err)
	}

	signed, err := jwt.Sign(token, jwt.WithKey(jwa.RS256, j.PrivateKey))
	if err != nil {
		return "", nil, fmt.Errorf("failed to sign token: %w", err)
	}

	return token.JwtID(), signed, nil
}

func (j *JWTer) ParseRequest(r *http.Request) (jwt.Token, error) {
	token, err := jwt.ParseRequest(
		r,
		jwt.WithKey(jwa.RS256, j.PublicKey),
		jwt.WithValidate(false))
	if err != nil {
		return nil, err
	}

	nowfunc := jwt.WithClock(jwt.ClockFunc(func() time.Time { return time.Now() }))
	if err := jwt.Validate(token, nowfunc); err != nil {
		return nil, fmt.Errorf("failed to validate token: %w", err)
	}

	return token, nil
}
