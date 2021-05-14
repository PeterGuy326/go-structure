package BST

import "fmt"

// 新树
func BSTNew(rootVal int) *Node {
	// 如果不写全，要指定 Key
	return &Node{value: rootVal}
}

// 插入树 (把节点看成一个简单的树)
func (bst *Node) Insert(val int) *Node {
	if bst == nil {
		return BSTNew(val)
	} else {
		if val < bst.value {
			bst.left = bst.left.Insert(val)
		} else {
			bst.right = bst.right.Insert(val)
		}
		return bst
	}
}

// 先序遍历
func (bst *Node) PreOrder() {
	if bst == nil {
		return
	} else {
		fmt.Printf("%d", bst.value)
		bst.left.PreOrder()
		bst.right.PreOrder()
	}
}

// 中序遍历（从小到大排序）
func (bst *Node) InOrder() {
	if bst == nil {
		return
	} else {
		bst.left.InOrder()
		fmt.Printf("%d", bst.value)
		bst.right.InOrder()
	}
}

//后序遍历
func (bst *Node) PostOrder() {
	if bst == nil {
		return
	} else {
		bst.left.PostOrder()
		bst.right.PostOrder()
		fmt.Printf("%d", bst.value)
	}
}
