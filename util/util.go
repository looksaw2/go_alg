package util

import (
	"math/rand"
	"time"
)

func GenArr(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	return arr
}

func GenArrWithSameValue[T any](n int, value T) []T {
	arr := make([]T, n)
	for i := 0; i < n; i++ {
		arr[i] = value
	}
	return arr
}

func GenArrRandom(n int) []int {
	src := rand.NewSource(time.Now().UnixMicro())
	r := rand.New(src)
	var arr []int
	for i := 0; i < n; i++ {
		arr = append(arr, r.Intn(10))
	}
	return arr
}
