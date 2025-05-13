package leetcode

import (
	"go_alg/util"
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func MiddleNode(head *ListNode) *ListNode {
	//快慢指针

	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func TrainingPlan(head *ListNode, cnt int) *ListNode {
	slow := head
	fast := head
	for i := 0; i < cnt; i++ {
		fast = fast.Next
	}
	for ; fast != nil; fast = fast.Next {
		slow = slow.Next
	}
	return slow
}

func TwoSum(price []int, target int) []int {
	if len(price) == 1 {
		return []int{}
	}
	N := len(price)
	left := 0
	right := N - 1
	for left <= right {
		if price[left]+price[right] == target {
			return []int{price[left], price[right]}
		} else if price[left]+price[right] > target {
			right--
		} else {
			left++
		}
	}
	return []int{}
}

func ClimbStairs(n int) int {
	// 1个台阶1种情况，2个台阶两个情况
	if n == 1 || n == 2 {
		return n
	}
	a, b := 1, 2
	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func MinCostClimbingStairs(costs []int) int {
	N := len(costs)
	if N == 1 {
		return costs[0]
	}
	if N == 2 {
		return costs[1]
	}
	dp := make([]int, N+1)
	dp[0] = costs[0]
	dp[1] = costs[1]
	for i := 2; i < N; i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + costs[i]
	}
	dp[N] = min(dp[N-1], dp[N-2])
	return dp[N]
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	mem := util.MakeSquareWithValue(m, n, -1)
	ans := pathSum(m-1, n-1, grid, mem)
	return ans
}

func pathSum(i int, j int, grid [][]int, mem [][]int) int {
	if i == 0 && j == 0 {
		return grid[0][0]
	}
	if i < 0 || j < 0 {
		return math.MaxInt
	}
	if mem[i][j] != -1 {
		return mem[i][j]
	}
	up := pathSum(i-1, j, grid, mem)
	left := pathSum(i, j-1, grid, mem)
	mem[i][j] = min(up, left) + grid[i][j]
	return mem[i][j]
}
