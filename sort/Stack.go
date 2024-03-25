package sortPkg

type StackElement[T any] struct {
	Value T
	Prev  *StackElement[T]
}

type Stack[T any] struct {
	Size int
	Top  *StackElement[T]
}

func (s *Stack[T]) Push(item T) {
	node := StackElement[T]{Value: item}
	s.Size++

	if s.Top == nil {
		s.Top = &node
		return
	}

	node.Prev = s.Top
	s.Top = &node
}

func (s *Stack[T]) Pop() *T {
	if s.Top == nil {
		return nil
	}

	s.Size--
	top := s.Top.Value
	s.Top = s.Top.Prev
	return &top
}

func (s *Stack[T]) Peek() *T {
	return &s.Top.Value
}
