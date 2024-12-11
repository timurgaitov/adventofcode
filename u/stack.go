package u

type stack struct {
	arr []int
	cur int
}

func NewStack() *stack {
	return &stack{
		arr: make([]int, 10000),
		cur: -1,
	}
}

func (s *stack) Empty() bool {
	return s.cur == -1
}

func (s *stack) Add(val int) {
	s.cur++
	s.arr[s.cur] = val
}

func (s *stack) Remove() int {
	if s.cur < 0 {
		panic("empty stack")
	}
	val := s.arr[s.cur]
	s.cur--
	return val
}

func (s *stack) Cur() int {
	if s.cur < 0 {
		panic("empty stack")
	}
	return s.arr[s.cur]
}
