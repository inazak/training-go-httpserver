package ttlmap

import (
	"testing"
	"time"
)

func Test1(t *testing.T) {
	t.Parallel()

	ps := map[string]int{
		"Hydrogen": 1,
		"Helium":   2,
		"Lithium":  3,
	}

	ttlmap := NewTTLMap()
	for k, v := range ps {
		ttlmap.Set(k, v, 60)
	}

	for k, v := range ps {
		r, ok := ttlmap.Get(k)
		if !ok || r != v {
			t.Errorf("key not found or value do not match, key=%s", k)
		}
	}
}

func Test2(t *testing.T) {
	t.Parallel()

	ps := map[string]int{
		"Hydrogen": 1,
		"Helium":   2,
		"Lithium":  3,
	}

	ttlmap := NewTTLMap()
	for k, v := range ps {
		ttlmap.Set(k, v, 2)
	}

	time.Sleep(time.Second * 3)

	for k, _ := range ps {
		_, ok := ttlmap.Get(k)
		if ok {
			t.Errorf("key has not expired, key=%s", k)
		}
	}
}

func Test3(t *testing.T) {
	t.Parallel()

	ttlmap := NewTTLMap()
	ttlmap.Set("Hydrogen", 1, 0)
	ttlmap.Set("Helium", 2, 1)
	ttlmap.Set("Lithium", 3, 1)

	time.Sleep(time.Second * 2)

	_, ok := ttlmap.Get("Hydrogen")
	if !ok {
		t.Errorf("key has expired, key=Hydrogen")
	}
	ttlmap.Delete("Hydrogen")

	_, ok = ttlmap.Get("Hydrogen")
	if ok {
		t.Errorf("key has not expired after delete, key=Hydrogen")
	}

	_, ok = ttlmap.Get("Helium")
	if ok {
		t.Errorf("key has not expired, key=Helium")
	}

	_, ok = ttlmap.Get("Lithium")
	if ok {
		t.Errorf("key has not expired, key=Lithium")
	}
}

func Test4(t *testing.T) {
	t.Parallel()

	ttlmap := NewTTLMap()
	ttlmap.Set("Hydrogen", 1, 2)
	ttlmap.Set("Hydrogen", 1, 60) //reset ttl

	time.Sleep(time.Second * 3)

	_, ok := ttlmap.Get("Hydrogen")
	if !ok {
		t.Errorf("key has expired, key=Hydrogen")
	}
}
