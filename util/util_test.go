package util_test

import (
	"go_alg/util"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	N = 10
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestGenArr(t *testing.T) {
	arr := util.GenArr(N)
	arr2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	require.Equal(t, arr, arr2, "arr应该和arr2相等")
}

func TestGenArrWithSameValue(t *testing.T) {
	arr := util.GenArrWithSameValue(N, 1)
	arr2 := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	require.Equal(t, arr, arr2, "arr1 should be eqaul to arr2")
}
