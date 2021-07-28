package RB

// getGrandParent() 获取父级节点的父级节点
func (rbnode *RBNode) getGrandParent() *RBNode {
	return rbnode.Parent
}

// getSibling() 获取兄弟节点
func (rbnode *RBNode) getSibling() *RBNode {
	return nil
}

// GetUncle() 父节点的兄弟节点
func (rbnode *RBNode) getUncle() *RBNode {
	return nil
}
