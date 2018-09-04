package safemap

type SafeMapInterface interface {
	GetObject(key string) (interface{}, bool)
	SetObject(key string, obj interface{}) error
	RemoveObject(key string) error
	Values() []interface{}
	Keys() []string
}
