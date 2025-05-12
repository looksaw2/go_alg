package sort

import (
	"golang.org/x/exp/constraints"
)

type SimpleSort[T constraints.Ordered] struct {
	DefaultSort[T]
}

func (s *SimpleSort[T]) Sort(arr []T) {
	for i := 0; i < len(arr); i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if s.DefaultSort.Less(arr[j], arr[minIndex]) {
				minIndex = j
			}
		}
		s.DefaultSort.Exch(arr, i, minIndex)
	}
}

type InsertionSort[T constraints.Ordered] struct {
	DefaultSort[T]
}

func (s *InsertionSort[T]) Sort(arr []T) {
	for i := 0; i < len(arr); i++ {
		for j := i; j >= 1 && s.DefaultSort.Less(arr[j], arr[j-1]); j-- {
			s.DefaultSort.Exch(arr, j, j-1)
		}
	}
}

type QuickSort[T constraints.Ordered] struct {
	DefaultSort[T]
}

func (s *QuickSort[T]) Sort(arr []T) {
	s.sort(arr, 0, len(arr)-1)
}

func (s *QuickSort[T]) sort(arr []T, lo int, hi int) {
	if hi <= lo {
		return
	}
	lt := lo
	gt := hi
	i := lt + 1
	v := arr[lt]
	for i <= gt {
		if arr[i] < v {
			s.DefaultSort.Exch(arr, lt, i)
			lt++
			i++
		} else if arr[i] > v {
			s.DefaultSort.Exch(arr, i, gt)
			gt--
		} else {
			i++
		}
	}
	s.sort(arr, lo, lt-1)
	s.sort(arr, gt+1, hi)
}
