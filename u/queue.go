package u

type queue struct {
	arr []int
	deq int
	enq int
}

func NewQueue() *queue {
	return &queue{
		arr: make([]int, 10000),
		deq: 0,
		enq: 0,
	}
}

func (s *queue) Empty() bool {
	return s.deq == s.enq
}

func (s *queue) Add(val int) {
	s.arr[s.enq] = val
	s.enq = (s.enq + 1) % len(s.arr)
	if s.enq == s.deq {
		panic("queue is full")
	}
}

func (s *queue) Remove() int {
	if s.enq == s.deq {
		panic("empty queue")
	}
	val := s.arr[s.deq]
	s.deq = (s.deq + 1) % len(s.arr)
	return val
}

func (s *queue) Cur() int {
	if s.enq == s.deq {
		panic("empty queue")
	}
	return s.arr[s.deq]
}
