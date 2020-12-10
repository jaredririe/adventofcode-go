package main

func FindInvalidNumber(input []int, preamble int) int {

	for i := preamble; i < len(input); i++ {

		target := input[i]

		if twoSum(input[i-preamble:i], target) == -1 {
			return input[i]
		}
	}

	return -1
}

func FindEncryptionWeakness(input []int, preamble int) int {

	invalid := FindInvalidNumber(input, preamble)
	if invalid == -1 {
		return -1
	}

	for i := 0; i < len(input); i++ {

		for j := i + 1; j < len(input); j++ {

			if sum(input[i:j+1]) == invalid {
				min := min(input[i : j+1])
				max := max(input[i : j+1])
				return min + max
			}
		}
	}

	return -1
}
