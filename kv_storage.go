package main

type KVStorage struct {
	storage map[string]Value
}

type Entry struct {
	key   string
	value Value
}

func NewKVStorage() *KVStorage {
	return &KVStorage{storage: make(map[string]Value)}
}

func (kv *KVStorage) Update(key string, value Value) bool {
	curr_value, ok := kv.storage[key]
	if !ok {
		kv.storage[key] = value
		return true
	}
	if Less(curr_value.tp, value.tp) {
		kv.storage[key] = value
		return true
	}
	return false
}

func (kv *KVStorage) Get(key string) (Value, bool) {
	value, ok := kv.storage[key]
	return value, ok
}
