package maxpq

import (
	"golang.org/x/exp/constraints"
)

type MaxPQ[Key constraints.Ordered] interface {
	Insert(v Key)
	Max() Key
	DelMax() Key
	IsEmpty() bool
	Size() int
}

type MyMqxPQ[Key constraints.Ordered] struct {
	pq []Key
	N  int
}

func NewMyMaxPQ[Key constraints.Ordered](maxN int) MyMqxPQ[Key] {
	return MyMqxPQ[Key]{
		pq: make([]Key, maxN+1),
		N:  0,
	}
}

func (s *MyMqxPQ[Key]) IsEmpty() bool {
	return s.N == 0
}

func (s *MyMqxPQ[Key]) Size() int {
	return s.N
}

func (s *MyMqxPQ[Key]) Insert(v Key) {
	s.N++
	s.pq[s.N] = v
	s.swim(s.N)
}

func (s *MyMqxPQ[Key]) DelMax() Key {
	max := s.pq[1]
	s.exch(1, s.N)
	s.N--
	s.sink(1)
	return max
}

func (s *MyMqxPQ[Key]) swim(k int) {
	for k > 1 && s.less(k/2, k) {
		s.exch(k/2, k)
	}
}

func (s *MyMqxPQ[Key]) sink(k int) {
	for 2*k <= s.N {
		j := 2 * k
		if j < s.N && s.less(j, j+1) {
			j++
		}
		if !s.less(k, j) {
			break
		}
		s.exch(k, j)
		k = j
	}
}

func (s *MyMqxPQ[Key]) less(i int, j int) bool {
	if s.pq[i] < s.pq[j] {
		return true
	}
	return false
}

func (s *MyMqxPQ[Key]) exch(i int, j int) {
	temp := s.pq[i]
	s.pq[i] = s.pq[j]
	s.pq[j] = temp
}
