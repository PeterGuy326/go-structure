package RB

const (
	// RED 红树设为true
	RED bool = true
	// BLACK 黑树设为false
	BLACK bool = false
)

const (
	// 左旋
	LEFTROTATE bool = true
	// 右旋
	RIGHTROTATE bool = false
)

// RBNode 红黑树
type RBNode struct {
	Value  int64
	Color  bool
	Left   *RBNode
	Right  *RBNode
	Parent *RBNode
}

type RBTree struct {
	Root *RBNode
}
