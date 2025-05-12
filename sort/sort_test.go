package sort

import (
	"github.com/stretchr/testify/require"
	"go_alg/util"
	"os"
	"testing"
)

const (
	N = 10
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestSimpleSort1(t *testing.T) {
	arr := util.GenArr(N)
	sort := SimpleSort[int]{}
	sort.Sort(arr)
	expect := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	require.Equal(t, arr, expect)
}

func TestSimpleSort2(t *testing.T) {
	arr := []int{5, 4, 1, 2, 3}
	sort := SimpleSort[int]{}
	sort.Sort(arr)
	expect := []int{1, 2, 3, 4, 5}
	require.Equal(t, arr, expect)
}

func TestSimpleSort3(t *testing.T) {
	arr := []int{1}
	sort := SimpleSort[int]{}
	sort.Sort(arr)
	expect := []int{1}
	require.Equal(t, arr, expect)
}

func TestInsertionSort1(t *testing.T) {
	arr := util.GenArr(N)
	sort := InsertionSort[int]{}
	sort.Sort(arr)
	expect := util.GenArr(N)
	require.Equal(t, arr, expect)
}

func TestInsertionSort2(t *testing.T) {
	arr := []int{5, 4, 1, 2, 3}
	sort := InsertionSort[int]{}
	sort.Sort(arr)
	expect := []int{1, 2, 3, 4, 5}
	require.Equal(t, arr, expect)
}

func TestInsertionSort3(t *testing.T) {
	arr := []int{1}
	sort := InsertionSort[int]{}
	sort.Sort(arr)
	expect := []int{1}
	require.Equal(t, arr, expect)
}

func TestQuickSort1(t *testing.T) {
	arr := []int{5, 4, 1, 2, 3}
	sort := QuickSort[int]{}
	sort.Sort(arr)
	expect := []int{1, 2, 3, 4, 5}
	require.Equal(t, arr, expect)
}
