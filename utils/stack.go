package utils

type stack struct {
	arr []any
	cur int
}

func NewStack() *stack {
	return &stack{
		arr: make([]any, 10000000000),
		cur: -1,
	}
}

func (s *stack) Empty() bool {
	return s.cur == -1
}

func (s *stack) Add(val any) {
	s.cur++
	s.arr[s.cur] = val
}

func (s *stack) Remove() any {
	if s.cur < 0 {
		panic("empty stack")
	}
	val := s.arr[s.cur]
	s.cur--
	return val
}

func (s *stack) Cur() any {
	if s.cur < 0 {
		panic("empty stack")
	}
	return s.arr[s.cur]
}
