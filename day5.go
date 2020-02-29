package main

func twoSum(nums []int, target int) []int {
	map1 := make(map[int]int)

	for index1, num := range nums {
		other := target - num

		index2, exist := map1[other]
		if exist {
			return []int{index2, index1}
		}
		map1[num] = index1
	}

	return []int{}
}