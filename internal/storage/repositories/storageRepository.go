package repositories

type StorageRepository interface {
	Get(key string) (value interface{}, ok bool)
	Set(key string, value interface{})
	GetAll() (data map[string]interface{})
}
