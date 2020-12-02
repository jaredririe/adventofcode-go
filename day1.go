package main

import "fmt"

func TwoSum1(input []int) int {

	target := 2020

	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			if input[i]+input[j] == target {
				fmt.Println(input[i], input[j])
				return input[i] * input[j]
			}
		}
	}

	return -1
}

func TwoSum2(input []int) int {

	target := 2020

	hash := make(map[int]int)

	for i1, v1 := range input {

		hash[v1] = i1

		complement := target - v1
		if i2, ok := hash[complement]; ok {
			fmt.Println(v1, input[i2])
			return v1 * input[i2]
		}
	}

	return -1
}

func ThreeSum(input []int) int {
	return findThreeSum(input, 0, 2020, []int{})
}

// findThreeSum is a recursive solution with a poor time complexity of O(2^N).
// As a positive, it can easily be adapted to find ANY length of combinations.
func findThreeSum(input []int, idx, target int, selections []int) int {

	// Base case: target has been found!
	if target == 0 && len(selections) == 3 {
		fmt.Println(selections)
		return selections[0] * selections[1] * selections[2]
	}

	// Base case: target is negative; as all inputs are positive, we can stop this branch
	if target < 0 {
		return -1
	}

	// Base case: past end of inputs slice
	if idx >= len(input) {
		return -1
	}

	// Base case: all selections have already been made
	if len(selections) >= 3 {
		return -1
	}

	// s1: current selections + current input
	// s2: current selections
	var s1, s2 []int
	for _, v := range selections {
		s1 = append(s1, v)
		s2 = append(s2, v)
	}

	s1 = append(s1, input[idx])

	if product := findThreeSum(input, idx+1, target-input[idx], s1); product != -1 {
		return product
	}

	if product := findThreeSum(input, idx+1, target, s2); product != -1 {
		return product
	}

	return -1
}
