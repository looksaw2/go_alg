package leetcode

import "sort"

func getMinLength(a []int, b []int) int {
	n := len(a)
	m := len(b)
	if n < m {
		return n
	} else {
		return m
	}
}

func isEnoughToFeed(need int, feed int) bool {
	return need <= feed
}

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	ans := 0
	numofchild := len(g)
	numofCook := len(s)
	for i, j := 0, 0; i < numofchild && j < numofCook; i++ {
		if isEnoughToFeed(g[i], s[j]) {
			j++
			ans++
			continue
		} else {
			sum := 0
			for ; !isEnoughToFeed(g[i], sum) && j < numofCook; j++ {
				sum += s[j]
			}
			if isEnoughToFeed(g[i], sum) {
				ans++
			} else {
				return ans
			}
		}
	}
	return ans
}

func isTop(s []int, i int) bool {
	left := 1
	for s[i-left] == -1 {
		left++
	}
	right := 1
	for s[i+right] == -1 {
		right++
	}
	if s[i] > s[i-left] && s[i] > s[i+right] {
		return true
	}
	return false

}
func isBottom(s []int, i int) bool {
	left := 1
	for s[i-left] == -1 {
		left++
	}
	right := 1
	for s[i+right] == -1 {
		right++
	}
	if s[i] < s[i-left] && s[i] < s[i+right] {
		return true
	}
	return false
}

func shouldNoEqual(nums []int) []int {
	var ans []int
	ans = append(ans, nums[0])
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue
		}
		ans = append(ans, nums[i])
	}
	return ans
}

func hasEqual(s []int, i int) bool {
	if s[i] == s[i-1] {
		return true
	}
	if s[i] == s[i+1] {
		return true
	}
	return false
}

func wiggleMaxLength(nums []int) int {
	return 0
}

func dailyTemperatures(temperatures []int) []int {
	var stack []int
	n := len(temperatures)
	//存index
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		toAdd := temperatures[i]
		for j := len(stack) - 1; j >= 0; j-- {
			//栈顶元素index
			index := stack[len(stack)-1]
			if toAdd > temperatures[index] {
				ans[index] = i - index
				stack = stack[:len(stack)-1]
			} else {
				break
			}
		}
		stack = append(stack, i)
	}
	return ans
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var stack []int
	n1 := len(nums1)
	n2 := len(nums2)

	temp := make(map[int]int)

	for i := 0; i < n2; i++ {
		toAdd := nums2[i]
		for j := len(stack) - 1; j >= 0; j-- {
			if toAdd > stack[j] {
				temp[stack[j]] = toAdd
				stack = stack[:len(stack)-1]
			} else {
				break
			}
		}
		stack = append(stack, toAdd)
	}
	var ans []int
	for i := 0; i < n1; i++ {
		v := temp[nums1[i]]
		if v == 0 {
			ans = append(ans, -1)
		} else {
			ans = append(ans, v)
		}
	}
	return ans
}
