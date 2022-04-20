package main

import (
	"fmt"
	"sort"
)

func main() {
	raw := []int{1,2,-2,-1}
	res := threeSum(raw)
	fmt.Printf("res:%v", res)
}


// 求三和为0
func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	res := make([][]int, 0)

	for first:=0; first<n; first++{
		if first>=1 && nums[first] == nums[first-1] {
			continue
		}

		third := n-1

		for second:=first+1; second<n;second++{
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}

			target := -1*(nums[first] + nums[second])

			for second<third && nums[third] > target {
				third--
			}

			if second >= third {
				break
			}

			if nums[third] == target {
				fmt.Printf("%d,%d,%d\n", first,second,third)
				res = append(res, []int{nums[first], nums[second], nums[third]})
			}
		}
	}

	return res
}


// 求最长公共前缀
func longestCommonPrefix(strs []string) string {
	// 条件已说明至少有一个元素
	prefix := strs[0]

	for i:=1; i<len(strs); i++ {
		prefix = maxCommon(prefix, strs[i])
		if prefix == "" {
			break
		}
	}

	return prefix
}
func maxCommon(prefix, compare string) string {
	n := len(prefix)
	if n > len(compare) {
		n = len(compare)
	}
	res := make([]byte, 0)

	for i:=0; i<n; i++ {
		if prefix[i] == compare[i] {
			res = append(res, prefix[i])
			continue
		}
		break
	}

	return string(res)
}