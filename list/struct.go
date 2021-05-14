package list

type List struct {
	Head   *ListNode
	length int
}

type ListNode struct {
	Next *ListNode
	Data interface{}
}
