package list

// 链表中添加节点
func (p *List) AddNode(data interface{}) {
	if p.Head == nil {
		p.Head = new(ListNode)
	}
	node := &ListNode{
		Data: data,
	}
	currentNext := p.Head.Next
	p.Head.Next = node
	node.Next = currentNext
}

// 获取链表获取长度
func (p *List) Length() int {
	return p.length
}

// 根据位置删除元素
func (p *List) DeleteWithPos(pos int) {
	if pos+1 > p.length {
		return
	}
	var i int
	pre := p.Head
	for {
		if pre.Next == nil {
			break
		}
		if i == pos {
			pre.Next = pre.Next.Next
		}
		pre = pre.Next
		i++
	}
	return
}

// 根据元素的值删除元素
func (p *List) DeleteWithData(value interface{}) {
	pre := p.Head
	for {
		if pre.Next == nil {
			break
		}
		if value == pre.Next.Data {
			pre.Next = pre.Next.Next
		}
		pre = pre.Next
	}
	return
}
