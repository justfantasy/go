package leetcode

func TwoSum(nums []int, target int) []int {
	tmp := make(map[int]int)
	for i, v := range nums {
		if _, ok := tmp[v]; ok {
			return []int{tmp[v], i}
		}
		tmp[target - v] = i
	}
	return nil
}