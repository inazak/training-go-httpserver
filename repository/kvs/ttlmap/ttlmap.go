package ttlmap

import (
	"sync"
	"time"
)

type record struct {
	Value      any
	Expiration int64
}

type TTLMap struct {
	*sync.RWMutex
	mapping map[string]record
}

func NewTTLMap() *TTLMap {
	return &TTLMap{
		RWMutex: &sync.RWMutex{},
		mapping: map[string]record{},
	}
}

// ttl が 0 以下の場合は、無期限保持とする
func (m *TTLMap) Set(k string, v any, ttl int64) {

	timestamp := int64(0)
	if ttl > 0 {
		timestamp = time.Now().Unix() + ttl
	}

	r := record{
		Value:      v,
		Expiration: timestamp,
	}

	m.Lock()
	m.mapping[k] = r
	m.Unlock()
}

func (m *TTLMap) Get(k string) (v any, ok bool) {
	m.RLock()
	r, ok := m.mapping[k]
	m.RUnlock()

	if ok {
		if r.Expiration == 0 {
			v = r.Value
		} else if r.Expiration > time.Now().Unix() {
			v = r.Value
		} else {
			m.Delete(k)
			ok = false
		}
	}

	return v, ok
}

func (m *TTLMap) Delete(k string) {
	m.Lock()
	delete(m.mapping, k)
	m.Unlock()
}

