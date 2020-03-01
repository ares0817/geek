package main

func plusOne(digits []int) []int {

	l := len(digits)
	i := 0
	for i < l {
		if digits[i] != 9 {
			break
		}
		i++
	}

	//all num is 9
	if i == l {
		//add a zero
		digits = append(digits, 0)
		//first is 1
		digits[0] = 1
		//other is 0
		for i = 1; i <= l; i++ {
			digits[i] = 0
		}
		return digits
	}

	i = l - 1
	for i >= 0 {
		digits[i] = digits[i] + 1
		if digits[i] < 10 {
			return digits
		}
		digits[i] = 0
		i--
	}

	return digits
}
