package BST

import (
	"errors"
	"fmt"
)

// 新树
func New(rootVal, rootData string) *Node {
	// 如果不写全，要指定 Key
	return &Node{Value: rootVal, Data: rootData}
}

// 插入树 (把节点看成一个简单的树)
func (bst *Node) Insert(value, data string) error {
	// 如果树为nil，则抛出异常
	if bst == nil {
		bst = New(value, data)
		return nil
	}

	switch {
	// 如果插入的数值在树中早已存在，返回
	case value == bst.Value:
		return nil
	// 若插入的数值小于当前的value
	case value < bst.Value:
		if bst.Left == nil {
			bst.Left = &Node{Value: value, Data: data}
			return nil
		}
		return bst.Left.Insert(value, data)
	// 若插入的数值大于当前的value
	case value > bst.Value:
		if bst.Right == nil {
			bst.Right = &Node{Value: value, Data: data}
			return nil
		}
		return bst.Right.Insert(value, data)
	}
	return nil
}

// 先序遍历
func (bst *Node) PreOrder() {
	if bst == nil {
		return
	} else {
		fmt.Printf("%d", bst.Value)
		bst.Left.PreOrder()
		bst.Right.PreOrder()
	}
}

// 中序遍历（从小到大排序）
func (bst *Node) InOrder() {
	if bst == nil {
		return
	} else {
		bst.Left.InOrder()
		fmt.Printf("%d", bst.Value)
		bst.Right.InOrder()
	}
}

// 后序遍历
func (bst *Node) PostOrder() {
	if bst == nil {
		return
	} else {
		bst.Left.PostOrder()
		bst.Right.PostOrder()
		fmt.Printf("%d", bst.Value)
	}
}

// 查询
func (bst *Node) Find(s string) (string, bool) {
	if bst == nil {
		return "", false
	}
	switch {
	case s == bst.Value:
		return bst.Data, true
	case s < bst.Value:
		return bst.Left.Find(s)
	default:
		return bst.Right.Find(s)
	}
}

// 查询最大值
func (bst *Node) findMax(parent *Node) (*Node, *Node) {
	if bst.Right == nil {
		return bst, parent
	}
	return bst.Right.findMax(bst)
}

// replaceNode用来替换，本身由parent指向n节点的指针，使得parent指向replacement节点
func (bst *Node) replaceNode(parent, replacement *Node) error {
	if bst == nil {
		return errors.New("replaceNode() not allowed on a nil node")
	}

	if bst == parent.Left {
		parent.Left = replacement
		return nil
	}
	parent.Right = replacement
	return nil
}

func (bst *Node) Delete(s string, parent *Node) error {
	// 如果树为nil，抛出异常
	if bst == nil {
		return errors.New("value to be deleted does not exist in the tree")
	}
	switch {
	// value小于当前节点的数值
	case s < bst.Value:
		return bst.Left.Delete(s, bst)
		// value大于当前节点的数值
	case s > bst.Value:
		return bst.Right.Delete(s, bst)
		// value与当前节点的数值相等
	default:
		// 如果要删除的节点是叶子节点
		if bst.Left == nil && bst.Right == nil {
			err := bst.replaceNode(parent, nil)
			if err != nil {
				return err
			}
			return nil
		}
		// 如果要删除的节点左孩子节点为nil
		if bst.Left == nil {
			err := bst.replaceNode(parent, bst.Right)
			if err != nil {
				return err
			}
			return nil
		}
		// 如果要删除的节点右孩子节点为nil
		if bst.Right == nil {
			err := bst.replaceNode(parent, bst.Left)
			if err != nil {
				return err
			}
			return nil
		}
		// 如果有两个孩子节点，在左子树中找出最大的值
		replacement, replParent := bst.Left.findMax(bst)
		// 用replacement的数值替代n节点的值域
		bst.Value = replacement.Value
		bst.Data = replacement.Data
		// 删除这个替代节点
		return replacement.Delete(replacement.Value, replParent)
	}
}
