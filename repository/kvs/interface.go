package kvs

type KVS interface {
	Set(key string, v any, ttl int64)
	Get(key string) (any, bool)
	Delete(key string)
}
