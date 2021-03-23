package leetcode

//来源https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0003.Longest-Substring-Without-Repeating-Characters/
func LengthOfLongestSubstring_(s string) int {
	length := len(s)
	if length == 0 {
		return 0
	}

	var bitSet [256]bool
	result, left, right := 0, 0, 0
	for left < length {
		if bitSet[s[right]] {
			bitSet[s[left]] = false
			left++
		} else {
			bitSet[s[right]] = true
			right++
		}

		if result < right-left {
			result = right - left
		}

		if left + result >= length || right >= length {
			break
		}
	}
	return result
}

// 使用滑动窗口算法
func LengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var tmp = make(map[int32]int)  // 创建一个字典，存储的字符和其对应的位置
	longest, left, right := 0, 0, 0          // 最大长度，左边界，右边界
	for _, v := range s {
		right++
		if inx, ok := tmp[v]; ok {
			if inx > left {   // 忽略超出边界的元素
				left = inx    // 出现相同的元素，则将元素放到边界外，这样整个窗口中不存在相同的字符
			}
		}
		if longest < right - left {
			longest = right - left
		}
		if left + longest >= len(s) {
			break
		}
		tmp[v] = right
	}
	return longest
}
