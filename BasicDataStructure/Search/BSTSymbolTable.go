package Search

import (
	"BasicDataStructure/DataStructure"
	"fmt"
)

type BSTSymbolTable struct {
	root *DataStructure.BinaryTree
}

func NewBSTSymbolTable() *BSTSymbolTable {
	return &BSTSymbolTable{
		root: nil,
	}
}

func (this *BSTSymbolTable) Put(key, value interface{}) {
	this.root = put(this.root, key, value)
}
func put(cur *DataStructure.BinaryTree, key, value interface{}) *DataStructure.BinaryTree {
	if cur == nil {
		return DataStructure.NewBinaryTree(key, value)
	}
	switch key.(type) {
	case string:
		if cur.GetNode().GetKey().(string) > key.(string) {
			cur.SetLeft(put(cur.GetLeft(), key, value))
		} else if cur.GetNode().GetKey().(string) < key.(string) {
			cur.SetRight(put(cur.GetRight(), key, value))
		} else {
			cur.GetNode().SetValue(value)
		}
	case int:
		if cur.GetNode().GetKey().(int) > key.(int) {
			cur.SetLeft(put(cur.GetLeft(), key, value))
		} else if cur.GetNode().GetKey().(int) < key.(int) {
			cur.SetRight(put(cur.GetRight(), key, value))
		} else {
			cur.GetNode().SetValue(value)
		}
	case float64:
		if cur.GetNode().GetKey().(float64) > key.(float64) {
			cur.SetLeft(put(cur.GetLeft(), key, value))
		} else if cur.GetNode().GetKey().(float64) < key.(float64) {
			cur.SetRight(put(cur.GetRight(), key, value))
		} else {
			cur.GetNode().SetValue(value)
		}
	}
	cur.SetCount(size(cur.GetLeft()) + size(cur.GetRight()) + 1)
	return cur
}

func (this *BSTSymbolTable) Get(key interface{}) interface{} {
	if this.root == nil {
		return nil
	}
	return get(this.root, key)
}
func get(cur *DataStructure.BinaryTree, key interface{}) interface{} {
	if cur == nil {
		return nil
	}
	switch key.(type) {
	case string:
		if cur.GetNode().GetKey().(string) > key.(string) {
			return get(cur.GetLeft(), key)
		} else if cur.GetNode().GetKey().(string) < key.(string) {
			return get(cur.GetRight(), key)
		} else {
			return cur.GetNode().GetValue()
		}
	case int:
		if cur.GetNode().GetKey().(int) > key.(int) {
			return get(cur.GetLeft(), key)
		} else if cur.GetNode().GetKey().(int) < key.(int) {
			return get(cur.GetRight(), key)
		} else {
			return cur.GetNode().GetValue()
		}
	case float64:
		if cur.GetNode().GetKey().(float64) > key.(float64) {
			return get(cur.GetLeft(), key)
		} else if cur.GetNode().GetKey().(float64) < key.(float64) {
			return get(cur.GetRight(), key)
		} else {
			return cur.GetNode().GetValue()
		}
	}
	return nil
}
func (this *BSTSymbolTable) Delete(key interface{}) {
	this.root = deleteNode(this.root, key)
}

func deleteNode(cur *DataStructure.BinaryTree, key interface{}) *DataStructure.BinaryTree {
	if cur == nil {
		return nil
	}
	switch key.(type) {
	case string:
		if cur.GetNode().GetKey().(string) > key.(string) {
			cur.SetLeft(deleteNode(cur.GetLeft(), key))
		} else if cur.GetNode().GetKey().(string) < key.(string) {
			cur.SetRight(deleteNode(cur.GetRight(), key))
		} else {
			if cur.GetRight() == nil {
				return cur.GetLeft()
			}
			if cur.GetLeft() == nil {
				return cur.GetRight()
			}
			temp := cur
			cur = minKey(temp.GetRight())
			cur.SetRight(deleteMinKey(temp.GetRight()))
			cur.SetLeft(temp.GetLeft())
		}
	case int:
		if cur.GetNode().GetKey().(int) > key.(int) {
			cur.SetLeft(deleteNode(cur.GetLeft(), key))
		} else if cur.GetNode().GetKey().(int) < key.(int) {
			cur.SetRight(deleteNode(cur.GetRight(), key))
		} else {
			if cur.GetRight() == nil {
				return cur.GetLeft()
			}
			if cur.GetLeft() == nil {
				return cur.GetRight()
			}
			temp := cur
			cur = minKey(temp.GetRight())
			cur.SetRight(deleteMinKey(temp.GetRight()))
			cur.SetLeft(temp.GetLeft())
		}
	case float64:
		if cur.GetNode().GetKey().(float64) > key.(float64) {
			cur.SetLeft(deleteNode(cur.GetLeft(), key))
		} else if cur.GetNode().GetKey().(float64) < key.(float64) {
			cur.SetRight(deleteNode(cur.GetRight(), key))
		} else {
			if cur.GetRight() == nil {
				return cur.GetLeft()
			}
			if cur.GetLeft() == nil {
				return cur.GetRight()
			}
			temp := cur
			cur = minKey(temp.GetRight())
			cur.SetRight(deleteMinKey(temp.GetRight()))
			cur.SetLeft(temp.GetLeft())
		}
	}
	cur.SetCount(size(cur.GetLeft()) + size(cur.GetRight()) + 1)
	return cur
}

func (this *BSTSymbolTable) DeleteMinKey() {
	this.root = deleteMinKey(this.root)
}
func deleteMinKey(cur *DataStructure.BinaryTree) *DataStructure.BinaryTree {
	if cur.GetLeft() == nil {
		return cur.GetRight()
	}
	cur.SetLeft(deleteMinKey(cur.GetLeft()))
	cur.SetCount(size(cur.GetLeft()) + size(cur.GetRight()) + 1)
	return cur
}

func (this *BSTSymbolTable) Contains(key interface{}) bool {
	return get(this.root, key) != nil
}

func (this *BSTSymbolTable) IsEmpty() bool {
	return size(this.root) == 0
}

func (this *BSTSymbolTable) Size() int {
	if this.root != nil {
		return size(this.root)
	}
	return 0
}

func size(cur *DataStructure.BinaryTree) int {
	if cur != nil {
		return cur.GetCount()
	}
	return 0
}
func (this *BSTSymbolTable) KeySet() []interface{} {
	panic("implement me")
}

func (this *BSTSymbolTable) MinKey() interface{} {
	return minKey(this.root).GetNode().GetKey()
}
func minKey(cur *DataStructure.BinaryTree) *DataStructure.BinaryTree {
	if cur.GetLeft() == nil {
		return cur
	}
	return minKey(cur.GetLeft())
}

func (this *BSTSymbolTable) MaxKey() interface{} {
	return maxKey(this.root).GetNode().GetKey()
}
func maxKey(cur *DataStructure.BinaryTree) *DataStructure.BinaryTree {
	if cur.GetRight() == nil {
		return cur
	}
	return maxKey(cur.GetRight())
}

func (this *BSTSymbolTable) PrintMiddle() {
	printMiddle(this.root)
	fmt.Println()
}
func printMiddle(cur *DataStructure.BinaryTree) {
	if cur.GetLeft() != nil {
		printMiddle(cur.GetLeft())
	}
	fmt.Print(cur.GetNode().GetKey(), cur.GetNode().GetValue(), " ")
	if cur.GetRight() != nil {
		printMiddle(cur.GetRight())
	}
}
