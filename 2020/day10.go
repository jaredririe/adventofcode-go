package main

import (
	"fmt"
	"sort"
)

func ConnectAllAdapters(input []int) int {

	sort.Ints(input)

	var voltage int

	var oneDiff, threeDiff int
	threeDiff++ // your device is always three higher

	for i := 0; i < len(input); i++ {

		diff := input[i] - voltage

		if diff > 3 {
			return -1 // There's no adapter able to bridge the gap
		}

		if diff == 1 {
			oneDiff++
		} else if diff == 3 {
			threeDiff++
		}

		voltage += diff
	}

	return oneDiff * threeDiff
}

var table map[string]int

func CountValidCombinations(input []int) int {
	table = make(map[string]int)
	sort.Ints(input)
	return countValidCombinations(input, 0, 0, max(input))
}

func countValidCombinations(input []int, idx int, voltage, maxVoltage int) int {

	if voltage == maxVoltage {
		// This path led to a valid combination
		return 1
	}

	if idx > len(input)-1 {
		// We reached the end of the input
		return 0
	}

	// Check whether we've already explored this exact path before
	key := fmt.Sprintf("(%d, %d)", idx, voltage)
	if result, ok := table[key]; ok {
		return result
	}

	var count int

	diff := input[idx] - voltage
	if diff >= 1 && diff <= 3 {
		// Recursively explore a path that includes this idx's voltage
		count += countValidCombinations(input, idx+1, voltage+diff, maxVoltage)
	}

	// Recursively explore a path that does NOT include this idx's voltage
	count += countValidCombinations(input, idx+1, voltage, maxVoltage)

	// Cache the count
	table[key] = count

	return count
}
