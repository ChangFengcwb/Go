package Search

import "BasicDataStructure/DataStructure"

type linkedSymbolTable struct {
	head *DataStructure.LinkedList
	size int
}

func NewLinkedSymbolTable() *linkedSymbolTable {
	return &linkedSymbolTable{
		head: nil,
	}
}

func (this *linkedSymbolTable) Put(key, value interface{}) {
	if this.head == nil {
		this.head = DataStructure.NewLinkedList(key, value)
		this.size++
	} else {
		for cur := this.head; cur != nil; {
			if cur.GetKey() == key {
				cur.SetValue(value)
				return
			}
			if cur.Next == nil {
				cur.Next = DataStructure.NewLinkedList(key, value)
				this.size++
				return
			}
			cur = cur.Next
		}
	}
}

func (this *linkedSymbolTable) Get(key interface{}) interface{} {
	if this.head == nil {
		return nil
	}
	for cur := this.head; cur != nil; cur = cur.Next {
		if cur.GetKey() == key {
			return cur.GetValue()
		}
	}
	return nil
}

func (this *linkedSymbolTable) Delete(key interface{}) {
	this.Put(key, nil)
	this.size--
}
func (this *linkedSymbolTable) Contains(key interface{}) bool {
	return this.Get(key) != nil
}
func (this *linkedSymbolTable) IsEmpty() bool {
	return this.Size() == 0
}
func (this *linkedSymbolTable) Size() int {
	return this.size
}

func (this *linkedSymbolTable) KeySet() []interface{} {
	keys := make([]interface{}, this.size)
	for cur, i := this.head, 0; cur != nil; cur, i = cur.Next, i+1 {
		keys[i] = cur.GetKey()
	}
	return keys
}
