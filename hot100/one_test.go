package hot100_test

import (
	"go_alg/hot100"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGroupAnagrams(t *testing.T) {
	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	expected := [][]string{
		{"bat"},
		{"nat", "tan"},
		{"ate", "eat", "tea"},
	}
	output := hot100.GroupAnagrams(input)
	require.Equal(t, output, expected)
}

func TestLongestConsecutive(t *testing.T) {
	input := []int{1, 0, 1, 2}
	expected := 4
	output := hot100.LongestConsecutive(input)
	require.Equal(t, expected, output)
}

func TestThreeSum(t *testing.T) {
	input := []int{0, 0, 0}
	expected := [][]int{
		{0, 0, 0},
	}
	output := hot100.ThreeSum(input)
	require.Equal(t, output, expected)

}

func TestSubArraySum1(t *testing.T) {
	input := []int{1, 1, 1}
	expected := 2
	output := hot100.SubArraySum(input, 2)
	require.Equal(t, output, expected)

}

func TestMaxSlidingWindow(t *testing.T) {
	input := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	expected := []int{3, 3, 5, 5, 6, 7}
	output := hot100.MaxSlidingWindow(input, k)
	require.Equal(t, output, expected)
}

func TestMaxSubArray(t *testing.T) {
	input := []int{5, 4, -1, 7, 8}
	expected := 23
	output := hot100.MaxSubArray(input)
	require.Equal(t, output, expected)
}

func TestMaxSubArray2(t *testing.T) {
	input := []int{-2, 1}
	expected := 1
	output := hot100.MaxSubArray(input)
	require.Equal(t, output, expected)
}

func TestMerge(t *testing.T) {
	input := [][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}
	output := hot100.Merge(input)
	expected := [][]int{
		{1, 6},
		{8, 10},
		{15, 18},
	}
	require.Equal(t, output, expected)
}

func TestMerge1(t *testing.T) {
	input := [][]int{
		{1, 4},
		{4, 5},
	}
	output := hot100.Merge(input)
	expected := [][]int{
		{1, 5},
	}
	require.Equal(t, output, expected)
}

func TestMerge2(t *testing.T) {
	input := [][]int{
		{2, 3},
		{4, 5},
		{6, 7},
		{8, 9},
		{1, 10},
	}
	output := hot100.Merge(input)
	expected := [][]int{
		{1, 10},
	}
	require.Equal(t, output, expected)
}
