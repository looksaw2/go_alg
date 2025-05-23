package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMiddleNode1(t *testing.T) {
	head := ListNode{Val: 1}
	node2 := ListNode{Val: 2}
	node3 := ListNode{Val: 3}
	node4 := ListNode{Val: 4}
	node5 := ListNode{Val: 5}
	head.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5
	ans := MiddleNode(&head)
	require.Equal(t, ans.Val, 3)
}

func TestMiddleNode2(t *testing.T) {
	head := ListNode{Val: 1}
	node2 := ListNode{Val: 2}
	node3 := ListNode{Val: 3}
	node4 := ListNode{Val: 4}
	node5 := ListNode{Val: 5}
	node6 := ListNode{Val: 6}
	head.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5
	node5.Next = &node6
	ans := MiddleNode(&head)
	require.Equal(t, ans.Val, 4)
}

func TestTrainingPlan(t *testing.T) {
	head := ListNode{Val: 2}
	node2 := ListNode{Val: 4}
	node3 := ListNode{Val: 7}
	node4 := ListNode{Val: 8}
	head.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	ans := TrainingPlan(&head, 1)
	require.Equal(t, ans.Val, 8)
}

func TestTwoSum(t *testing.T) {
	arr := []int{3, 9, 12, 15}
	target := 18
	ans := TwoSum(arr, target)
	require.Equal(t, ans, []int{3, 15})
}

func TestClimbStairs(t *testing.T) {
	n := 4
	ans := ClimbStairs(n)
	require.Equal(t, ans, 5)
}

func TestMinCostClimbingStairs(t *testing.T) {
	cost := []int{10, 15, 20}
	ans := MinCostClimbingStairs(cost)
	require.Equal(t, ans, 15)
}

func TestMinCostClimbingStairs2(t *testing.T) {
	cost := []int{1, 100}
	ans := MinCostClimbingStairs(cost)
	require.Equal(t, ans, 1)
}

func TestFindContentChildren(t *testing.T) {
	g := []int{1, 2, 3}
	s := []int{1, 1}
	ans := findContentChildren(g, s)
	require.Equal(t, ans, 1)
}

func TestShouldEqual(t *testing.T) {
	arr1 := []int{2, 2, 2, 2, 2, 1}
	arr1 = shouldNoEqual(arr1)
	require.Equal(t, arr1, []int{2, 1})
	arr2 := []int{1, 2, 2, 2, 2, 2}
	arr2 = shouldNoEqual(arr2)
	require.Equal(t, arr2, []int{1, 2})
	arr3 := []int{2, 2, 2, 2, 2, 2}
	arr3 = shouldNoEqual(arr3)
	require.Equal(t, arr3, []int{2})
}

func TestDailyTemperatures(t *testing.T) {
	arr := []int{73, 74, 75, 71, 69, 72, 76, 73}
	ans := dailyTemperatures(arr)
	require.Equal(t, ans, []int{1, 1, 4, 2, 1, 1, 0, 0})
}

func TestNextGreaterElement(t *testing.T) {
	nums1 := []int{2, 4}
	nums2 := []int{1, 2, 3, 4}
	ans := nextGreaterElement(nums1, nums2)
	require.Equal(t, ans, []int{3, -1})
}
