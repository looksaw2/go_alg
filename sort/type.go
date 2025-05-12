package sort

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

type Sort[T constraints.Ordered] interface {
	Less(a T, b T) bool
	Exch(arr []T, i int, j int)
	Show(arr []T)
	IsSorted(arr []T) bool
}

type DefaultSort[T constraints.Ordered] struct{}

func (s *DefaultSort[T]) Less(a T, b T) bool {
	if a < b {
		return true
	}
	return false
}

func (s *DefaultSort[T]) Exch(arr []T, i int, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

func (s *DefaultSort[T]) Show(arr []T) {
	for i := 0; i < len(arr); i++ {
		fmt.Printf("arr[%d] is %v\n", i, arr[i])
	}
}

func (s *DefaultSort[T]) IsSorted(arr []T) bool {
	if len(arr) == 0 || len(arr) == 1 {
		return true
	}
	for i := 1; i < len(arr); i++ {
		if s.Less(arr[i], arr[i-1]) {
			return false
		}
	}
	return true
}
