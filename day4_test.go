package main

import "testing"

func TestMerge(t *testing.T) {
	nums1 := []int{1, 5, 8, 0, 0, 0, 0, 0}
	nums2 := []int{-1, 2, 2, 4, 6}
	m, n  := 3, 5
	merge(nums1, m, nums2, n)
}
