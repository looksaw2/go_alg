package hot100

import (
	"container/heap"
	"sort"
	"strings"
)

func MySortString(s string) string {
	chars := strings.Split(s, "")
	sort.Strings(chars)
	return strings.Join(chars, "")

}

func GroupAnagrams(strs []string) [][]string {
	HashMap := make(map[string][]string)
	for _, s := range strs {
		sSorted := MySortString(s)
		if _, ok := HashMap[sSorted]; ok {
			HashMap[sSorted] = append(HashMap[sSorted], s)
		} else {
			HashMap[sSorted] = append(HashMap[sSorted], s)
		}
	}
	var ans [][]string
	for _, value := range HashMap {
		ans = append(ans, value)
	}
	return ans
}
func getMax(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func LongestConsecutive(nums []int) int {
	sort.Ints(nums)
	max := 0
	ins := 0
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return 1
	}
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] == 1 {
			ins++
			max = getMax(max, ins)
		} else if nums[i] == nums[i-1] {
			continue
		} else {
			max = getMax(max, ins)
			ins = 0
		}
	}
	return max + 1
}

func isGreaterThanZero(a int, b int) bool {
	if a+b > 0 {
		return true
	}
	return false
}

func ThreeSumWithRange(index int, nums []int, ans *[][]int) {
	if index > 0 && nums[index] == nums[index-1] {
		return
	}
	left := index + 1
	right := len(nums) - 1
	for left < right {
		if nums[index]+nums[left]+nums[right] == 0 {
			*ans = append(*ans, []int{nums[index], nums[left], nums[right]})
			left++
			for left < right && nums[left] == nums[left-1] {
				left++
			}

			right--
			for left < right && nums[right] == nums[right+1] {
				right--
			}
		} else if nums[index]+nums[left]+nums[right] > 0 {
			right--
		} else if nums[index]+nums[left]+nums[right] < 0 {
			left++
		}
	}

}

func ThreeSum(nums []int) [][]int {
	n := len(nums)
	right := n - 1
	ans := make([][]int, 0)
	if n == 0 {
		return [][]int{}
	}
	if n == 1 {
		if nums[0] == 1 {
			a := []int{1}
			ans = append(ans, a)
			return ans
		}
	}
	sort.Ints(nums)
	for i := 0; i < right; i++ {
		ThreeSumWithRange(i, nums, &ans)
	}
	return ans
}

type NumArray struct {
	nums      []int
	prefixSum []int
}

func Constructor(nums []int) NumArray {
	prefixSum := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			prefixSum[i] = nums[i]
		} else {
			prefixSum[i] = prefixSum[i-1] + nums[i]
		}
	}
	return NumArray{
		nums:      nums,
		prefixSum: prefixSum,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.prefixSum[right]
	}
	return this.prefixSum[right] - this.prefixSum[left-1]
}

func getPrefixSum(nums []int) []int {
	prefixSum := make([]int, len(nums)+1)
	prefixSum[0] = 0
	for i := 1; i < len(nums)+1; i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i-1]

	}
	return prefixSum
}

func SubArraySum(nums []int, k int) int {
	prefixSum := 0
	prefizxSumMap := make(map[int]int)
	ans := 0
	prefizxSumMap[0] = 1
	for _, num := range nums {
		prefixSum += num
		if freq, exists := prefizxSumMap[prefixSum-k]; exists {
			ans += freq
		}
		prefizxSumMap[prefixSum]++
	}
	return ans
}

type A struct {
	index int
	value int
}

type AHeap []A

func (h AHeap) Len() int           { return len(h) }
func (h AHeap) Less(i, j int) bool { return h[i].value > h[j].value } // 大顶堆
func (h AHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *AHeap) Push(x any)        { *h = append(*h, x.(A)) }
func (h *AHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func MaxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}
	//新建一个大顶堆
	h := &AHeap{}
	heap.Init(h)
	ans := make([]int, 0)
	//新建一个Map key 为 index ， value 为是否需要删除
	toDeleteMap := make(map[int]bool)
	//0..k-1的初始化
	for i := 0; i < k; i++ {
		tA := A{
			index: i,
			value: nums[i],
		}
		heap.Push(h, tA)
	}
	//添加初始元素
	ans = append(ans, (*h)[0].value)
	//添加后续元素
	for i := k; i < len(nums); i++ {
		toDeleteMap[i-k] = true
		for h.Len() > 0 && toDeleteMap[(*h)[0].index] == true {
			toDeleteMap[(*h)[0].index] = false
			heap.Pop(h)
		}
		tA := A{
			index: i,
			value: nums[i],
		}
		heap.Push(h, tA)
		ans = append(ans, (*h)[0].value)
	}
	return ans
}

func MaxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	prefixSum := make([]int, len(nums)+1)
	prefixSum[0] = 0
	var min int
	var max int
	maxIndex := nums[0]
	minIndex := 0
	ans := 0
	for i := 1; i < len(nums)+1; i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i-1]
		if prefixSum[i] > max {
			max = prefixSum[i]
			maxIndex = i
		}
		if prefixSum[i] < min {
			min = prefixSum[i]
			minIndex = i
		}
		if maxIndex >= minIndex {
			ans = max - min
		}
	}
	return ans

}

func Merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i int, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var ans [][]int
	for i := 0; i < len(intervals); i++ {
		arr := intervals[i]
		for i < len(intervals)-1 && arr[1] >= intervals[i+1][0] {
			arr[1] = getMax(arr[1], intervals[i+1][1])
			i++
		}
		ans = append(ans, arr)
	}
	return ans
}

func Rotate(nums []int, k int) {
	aux := make([]int, 0)
	for i := 0; i < k; i++ {
		aux = append(aux, nums[len(nums)-1-i])
	}
	for i := len(nums) - 1 - k; i > 0; i-- {
		nums[i+k] = nums[i]
	}
	for i := 0; i < k; i++ {
		nums[i] = aux[k-1-i]
	}
}
