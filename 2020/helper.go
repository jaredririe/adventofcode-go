package main

func sum(input []int) int {

	var sum int
	for _, i := range input {
		sum += i
	}

	return sum
}

func min(input []int) int {

	smallest := input[0]
	for i := 1; i < len(input); i++ {
		if input[i] < smallest {
			smallest = input[i]
		}
	}

	return smallest
}

func max(input []int) int {

	largest := input[0]
	for i := 1; i < len(input); i++ {
		if input[i] > largest {
			largest = input[i]
		}
	}

	return largest
}

func twoSum(input []int, target int) int {

	hash := make(map[int]int)

	for i1, v1 := range input {

		hash[v1] = i1

		complement := target - v1
		if i2, ok := hash[complement]; ok {
			return v1 * input[i2]
		}
	}

	return -1
}
