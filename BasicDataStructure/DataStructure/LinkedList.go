package DataStructure

type LinkedList struct {
	key   interface{}
	value interface{}
	Next  *LinkedList
}

func NewLinkedList(key, value interface{}) *LinkedList {
	return &LinkedList{
		key:   key,
		value: value,
		Next:  nil,
	}
}

func (this *LinkedList) GetKey() interface{} {
	return this.key
}
func (this *LinkedList) SetKey(key interface{}) {
	this.key = key
}
func (this *LinkedList) GetValue() interface{} {
	return this.value
}
func (this *LinkedList) SetValue(value interface{}) {
	this.value = value
}
