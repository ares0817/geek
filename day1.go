package main

func removeDuplicates(nums []int) int {

	i, j, max := 0, 1, len(nums)

	for j < max {
		if nums[i] < nums[j] {
			i++
			//need swap
			if j > i {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}

		j++
	}

	//nums[i] is the last num, so return i + 1
	return i + 1
}