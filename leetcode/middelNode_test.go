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
