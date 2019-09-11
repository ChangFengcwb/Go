package DataStructure

type node struct {
	value interface{}
	next  *node
	pre   *node
}

func newNode(value interface{}) *node {
	return &node{
		value: value,
		next:  nil,
		pre:   nil,
	}
}

type keyValueNode struct {
	key   interface{}
	value interface{}
}

func NewKeyValueNode(key interface{}, value interface{}) *keyValueNode {
	return &keyValueNode{key: key, value: value}
}

func (this *keyValueNode) GetKey() interface{} {
	return this.key
}
func (this *keyValueNode) GetValue() interface{} {
	return this.value
}
func (this *keyValueNode) SetKey(key interface{}) {
	this.key = key
}
func (this *keyValueNode) SetValue(value interface{}) {
	this.value = value
}
