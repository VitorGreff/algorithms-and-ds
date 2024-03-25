package sortPkg

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type Queue[T any] struct {
	Lenght int
	Head   *Node[T]
	Tail   *Node[T]
}

func (q *Queue[T]) Enqueue(item T) {
	node := &Node[T]{Value: item}
	q.Lenght++

	if q.Tail == nil {
		q.Head = node
		q.Tail = node
		return
	}

	q.Tail.Next = node
	q.Tail = node
}

func (q *Queue[T]) Deque() *T {
	if q.Head == nil {
		return nil
	}

	q.Lenght--

	Value := q.Head.Value
	q.Head = q.Head.Next
	if q.Head == nil {
		q.Tail = nil
	}
	return &Value
}

func (q *Queue[T]) Peek() T {
	return q.Head.Value
}
