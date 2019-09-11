package Search

type SymbolTableApi interface {
	Put(key, value interface{})
	Get(key interface{}) interface{}
	Delete(key interface{})
	Contains(key interface{}) bool
	IsEmpty() bool
	Size() int
	KeySet() []interface{}
}
