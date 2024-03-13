package storage

var GlobalMemStorage *MemStorage

func init() {
	GlobalMemStorage = NewMemStorage()
}

type MemStorage struct {
	data map[string]interface{}
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		data: make(map[string]interface{}),
	}
}

func (m *MemStorage) Get(key string) (value interface{}, ok bool) {
	value, ok = m.data[key]
	return value, ok
}

func (m *MemStorage) Set(key string, value interface{}) {
	m.data[key] = value
}
