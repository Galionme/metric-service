package repositories

type UpdateRepository interface {
	Get(key string) (value interface{}, ok bool)
	Set(key string, value interface{})
}
