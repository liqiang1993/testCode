package _go

func longest(raw string) int {
	lastNum := 0
	currentNum := 0
	recordMap := make(map[byte]struct{})
	start := 0
	for index:= start;index < len(raw); index++ {
		if _, ok := recordMap[raw[index]]; ok {
			if currentNum > lastNum {
				lastNum = currentNum
			}
			start += 1
			index = start
			recordMap = make(map[byte]struct{})
			currentNum = 0
		}  else {
			recordMap[raw[index]] = struct{}{}
			currentNum += 1
		}
	}
	return lastNum
}


func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	total := len(nums1) + len(nums2)
	middleIndex := total/2
	twoFlag := false

	if total % 2 == 0 {
		twoFlag = true
	}

	index1 := 0
	index2 := 0
	sum := 0

	temp := 0
	for index := 0; index <= middleIndex; index++ {
		if nums1[index1] <= nums2[index2] {
			temp = nums1[index1]
			index1++
		} else {
			temp = nums2[index2]
			index2++
		}

		if twoFlag &&(index == middleIndex -1) {
			sum += temp
		}
	}

	sum += temp

	if twoFlag {
		return float64(sum) /2.0
	} else {
		return float64(sum)
	}
}
