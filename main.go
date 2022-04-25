package main

import (
	"fmt"
	"sort"
)

func main() {
	raw := [][]int{{1,1},{2,2},{3,4},{4,5},{5,6},{6,7}}
	res := checkStraightLine(raw)
	fmt.Printf("res:%v\n", res)
}

// 判断一组点是否在直线上
func checkStraightLine(coordinates [][]int) bool {
	base := float32(coordinates[1][1]- coordinates[0][1]) / float32(coordinates[1][0]- coordinates[0][0])
	n := len(coordinates)
	fmt.Printf("base:%v\n", base)
	for i:=2; i<n; i++ {
		temp := float32(coordinates[i][1]- coordinates[0][1]) / float32(coordinates[i][0]- coordinates[0][0])
		fmt.Printf("%v,%v,%v,%v\n", coordinates[i][1], coordinates[0][1], coordinates[i][0], coordinates[0][0])
		fmt.Printf("i:%v, temp:%v\n", i, temp)
		if temp != base {
			return false
		}
	}

	return true
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

//冒泡
func bubble(raws []int) {
	n := len(raws)
	for i:=n; i>1; i-- {
		for j:=1; j<n; j++ {
			if raws[j] < raws[j-1] {
				temp := raws[j]
				raws[j] = raws[j-1]
				raws[j-1] = temp
			}
		}
	}
}

//插入
func insert(nums []int) {
	n := len(nums)
	for i:=1; i<n; i++ {
		for j:=i; j>=1; j-- {
			if nums[j] < nums[j-1] {
				temp := nums[j]
				nums[j] = nums[j-1]
				nums[j-1] = temp
				continue
			}
			break
		}
	}
}


//选择
func choose(nums []int) {
	n := len(nums)

	for i:=0; i<n; i++ {
		index := i
		value := nums[i]
		for j:=i+1; j<n; j++ {
			if value > nums[j] {
				value = nums[j]
				index = j
			}
		}
		if i != index {
			temp := nums[i]
			nums[i] = value
			nums[index] = temp
		}
	}
}

//归并
func sortWithMerge(nums []int) {
	merge(nums, 0, len(nums)-1)
}

func merge(nums []int, left int, right int) {
	if left >= right {
		return
	}

	mid := left + (right-left)/2
	res := make([]int, 0, right-left+1)

	merge(nums, left, mid)
	merge(nums, mid+1, right)

	i:=left
	j:=mid+1

	for i<=mid&&j<=right {
		if nums[i] <= nums[j] {
			res = append(res, nums[i])
			i++
		} else {
			res = append(res, nums[j])
			j++
		}
	}

	for i<=mid {
		res = append(res, nums[i])
		i++
	}

	for j<=right {
		res = append(res, nums[j])
		j++
	}

	for index:=left; index<=right; index++ {
		nums[index] = res[index-left]
	}
}

//快速
func quickSort(nums []int) {
	subQuick(nums, 0, len(nums)-1)
}

func  subQuick(nums []int, left int, right int) {
	if left >= right {
		return
	}

	flagValue := nums[left]
	i := left
	j := right
	for i<j {
		for i<j && flagValue <= nums[j] {
			j--
		}
		nums[i] = nums[j]
		i++

		for i<j && flagValue >= nums[i] {
			i++
		}
		nums[j] = nums[i]
		j--
	}

	subQuick(nums, left, i-1)
	subQuick(nums, i+1, right)
}

//二分,must in
func quickFind(nums []int, target int, left int, right int) int {
	mid := left + (right-left)/2
	if target == nums[mid] {
		return mid
	}
    if target < nums[mid] {
    	return  quickFind(nums, target, left, mid-1)
	} else {
		return  quickFind(nums, target, mid+1, right)
	}
}




