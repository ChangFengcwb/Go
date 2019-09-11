package DataStructure

type BinaryTree struct {
	node  *keyValueNode
	count int
	left  *BinaryTree
	right *BinaryTree
}

func NewBinaryTree(key, value interface{}) *BinaryTree {
	return &BinaryTree{
		node:  NewKeyValueNode(key, value),
		count: 1,
		left:  nil,
		right: nil,
	}
}

func (this *BinaryTree) GetCount() int {
	return this.count
}
func (this *BinaryTree) GetNode() *keyValueNode {
	return this.node
}
func (this *BinaryTree) GetLeft() *BinaryTree {
	return this.left
}
func (this *BinaryTree) GetRight() *BinaryTree {
	return this.right
}
func (this *BinaryTree) SetCount(count int) {
	this.count = count
}
func (this *BinaryTree) SetNode(node *keyValueNode) {
	this.node = node
}
func (this *BinaryTree) SetLeft(tree *BinaryTree) {
	this.left = tree
}
func (this *BinaryTree) SetRight(tree *BinaryTree) {
	this.right = tree
}
