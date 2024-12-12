package utils

type queue struct {
	arr []any
	deq int
	enq int
}

func NewQueue() *queue {
	return &queue{
		arr: make([]any, 10000),
		deq: 0,
		enq: 0,
	}
}

func (s *queue) Empty() bool {
	return s.deq == s.enq
}

func (s *queue) Add(val any) {
	if s.next(s.enq) == s.deq {
		arr2 := make([]any, len(s.arr)*2)
		println("next size", len(arr2))
		copy(arr2, s.arr)
		s.arr = arr2
	}
	s.arr[s.enq] = val
	s.enq = s.next(s.enq)
}

func (s *queue) Remove() any {
	if s.enq == s.deq {
		panic("empty queue")
	}
	val := s.arr[s.deq]
	s.deq = s.next(s.deq)
	return val
}

func (s *queue) Cur() any {
	if s.enq == s.deq {
		panic("empty queue")
	}
	return s.arr[s.deq]
}

func (s *queue) next(cur int) int {
	return (cur + 1) % len(s.arr)
}
