package BST

// name:Binany Searching Tree 二叉搜索树
// description:二叉搜索树，是按照一个节点的左子叶小于节点的值，右子叶大于节点的值 搜索二叉树的，插入和搜索时间复杂都是O(lgn).

type Node struct {
	value int
	left  *Node
	right *Node
}
