package list

type (
	List interface {
		Add(val, priority int)
		Delete(n *Node)
		GetMin() *Node
		GetMax() *Node
	}

	list struct {
		head *Node
		tail *Node
	}

	Node struct {
		Val      int
		Priority int
		IsEnd    bool
		Prev     *Node
		Next     *Node
	}
)

func NewList() List {
	head := &Node{
		IsEnd: true,
	}
	head.Next = &Node{
		IsEnd: true,
		Prev:  head,
	}

	return &list{
		head: head,
		tail: head.Next,
	}
}

func (l *list) Add(val, priority int) {
	n := &Node{
		Val:      val,
		Priority: priority,
	}

	current := l.head
	for !current.Next.IsEnd && current.Next.Priority > n.Priority {
		current = current.Next
	}

	tmp := current.Next
	current.Next = n

	n.Prev = current
	n.Next = tmp

	tmp.Prev = n
}

func (l *list) Delete(n *Node) {
	if n.Prev != nil && !n.IsEnd {
		n.Prev.Next = n.Next
	}

	if n.Next != nil && !n.IsEnd {
		n.Next.Prev = n.Prev
	}
}

func (l *list) GetMin() *Node {
	if v := l.tail.Prev; !v.IsEnd {
		return v
	}
	return &Node{Val: -1}
}

func (l *list) GetMax() *Node {
	if v := l.head.Next; !v.IsEnd {
		return v
	}
	return &Node{Val: -1}
}
