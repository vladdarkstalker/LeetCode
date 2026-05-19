package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var result float64
	var generalLen = len(nums1) + len(nums2)
	var leftCount = generalLen / 2

	var buffer []int
	var i int
	var j int

	for true {
		if i == len(nums1) {
			buffer = append(buffer, nums2[j:]...)
			break
		} else if j == len(nums2) {
			buffer = append(buffer, nums1[i:]...)
			break
		}

		if nums1[i] < nums2[j] {
			buffer = append(buffer, nums1[i])
			i++
		} else {
			buffer = append(buffer, nums2[j])
			j++
		}
	}

	if generalLen%2 != 0 {
		result = float64(buffer[leftCount])
	} else {
		result = float64(buffer[leftCount-1]+buffer[leftCount]) / 2
	}

	return result
}
