package main

func merge(nums1 []int, m int, nums2 []int, n int)  {

	if n == 0 {
		return
	}
	i := m
	for _, num := range nums2 {
		nums1[i] = num
		i++
	}

	i, j := 0, m
	for i < j && nums1[i] <= nums1[j] {
		i++
	}

	for i < j {
		nums1[i], nums1[j] = nums1[j], nums1[i]
		for k := j; k < m + n - 1; k++ {
			if nums1[k] > nums1[k + 1] {
				nums1[k], nums1[k + 1] = nums1[k + 1], nums1[k]
			} else {
				break
			}
		}
		i++
	}
}