\# Intuition

Merge both sorted arrays into one sorted array, then take the middle element.



\# Approach

Use two pointers to compare elements from `nums1` and `nums2`.  

Append the smaller element to a new array.  

When one array ends, append the rest of the other array.  

After that, calculate the median depending on whether the total length is odd or even.



\# Complexity

\- Time complexity: $$O(n + m)$$



\- Space complexity: $$O(n + m)$$



\# Code

```golang \[]

func findMedianSortedArrays(nums1 \[]int, nums2 \[]int) float64 {

&#x20;   var result float64

&#x20;   var generalLen = len(nums1) + len(nums2)

&#x20;   var leftCount = generalLen / 2



&#x20;   var buffer \[]int

&#x20;   var i int

&#x20;   var j int



&#x20;   for true {

&#x20;       if i == len(nums1) {

&#x20;           buffer = append(buffer, nums2\[j:]...)

&#x20;           break

&#x20;       } else if j == len(nums2) {

&#x20;           buffer = append(buffer, nums1\[i:]...)

&#x20;           break

&#x20;       }



&#x20;       if nums1\[i] < nums2\[j] {

&#x20;           buffer = append(buffer, nums1\[i])

&#x20;           i++

&#x20;       } else {

&#x20;           buffer = append(buffer, nums2\[j])

&#x20;           j++

&#x20;       }

&#x20;   }



&#x20;   if generalLen%2 != 0 {

&#x20;       result = float64(buffer\[leftCount])

&#x20;   } else {

&#x20;       result = float64(buffer\[leftCount-1]+buffer\[leftCount]) / 2

&#x20;   }



&#x20;   return result

}

