package main

func rotate(nums []int, k int)  {

	l := len(nums)
	k = k % l
	if k == 0 {
		return
	}

	reverse(nums, 0, l - 1)
	reverse(nums, 0, k - 1)
	reverse(nums, k, l - 1)
}

func reverse(nums []int, start, end int) {

	i, j := start, end
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}