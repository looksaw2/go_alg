package uniofound

import util "go_alg/util"

type UF interface {
	Union(p int, q int)
	Find(p int) int
	Connected(p int, q int) bool
	Count() int
}

type MyUF struct {
	Id    []int
	sz    []int
	count int
}

func NewMyUF(n int) MyUF {
	arr := util.GenArr(n)
	return MyUF{
		Id:    arr,
		sz:    util.GenArrWithSameValue(n, 1),
		count: n,
	}
}

func (s *MyUF) Count() int {
	return s.count
}

func (s *MyUF) Connected(p int, q int) bool {
	return s.Find(p) == s.Find(q)
}

func (s *MyUF) Find(p int) int {
	return s.Id[p]
}

func (s *MyUF) Union(p int, q int) {
	pID := s.Find(p)
	qID := s.Find(q)
	if pID == qID {
		return
	}
	for i := 0; i < s.count; i++ {
		if s.Id[i] == qID {
			s.Id[i] = pID
		}
	}
	s.count--
}

func (s *MyUF) QuickFind(p int) int {
	for p != s.Id[p] {
		p = s.Id[p]
	}
	return p
}

func (s *MyUF) QuickUnion(p int, q int) {
	pID := s.QuickFind(p)
	qID := s.QuickFind(q)
	if pID == qID {
		return
	}
	if s.sz[pID] < s.sz[qID] {
		s.Id[pID] = qID
		s.sz[qID] += s.sz[pID]
	}
	s.count--
}
