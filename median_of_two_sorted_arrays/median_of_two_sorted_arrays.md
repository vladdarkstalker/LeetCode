# Intuition
Merge both sorted arrays into one sorted array, then take the middle element.

# Approach
Use two pointers to compare elements from `nums1` and `nums2`. Append the smaller element to a new array. When one array ends, append the rest of the other array. After that, calculate the median depending on whether the total length is odd or even.

# Complexity
- Time complexity: $$O(n + m)$$

- Space complexity: $$O(n + m)$$

# Code
```golang []
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var result float64
	var generalLen = len(nums1) + len(nums2)
	var leftCount = (generalLen / 2)

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
		result = float64((buffer[leftCount-1] + buffer[leftCount])) / 2
	}

	return result
}
```
