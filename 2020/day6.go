package main

func GroupCountsInclusive(input []string) int {

	groups := parseGroupAnswers(input)
	counts := make(map[rune]int)

	for _, group := range groups {

		// Find all unique choices in any answer in each group
		seen := make(map[rune]bool)
		for _, answers := range group {
			for _, choice := range answers {
				seen[choice] = true
			}
		}

		// Count the number of unique choices in the group
		for k := range seen {
			counts[k]++
		}
	}

	var groupCount int
	for _, v := range counts {
		groupCount += v
	}

	return groupCount
}

func GroupCountsExclusive(input []string) int {

	groups := parseGroupAnswers(input)
	counts := make(map[rune]int)

	for _, group := range groups {

		// Count the number of answers that included each choice
		// <choice> -> <count of answers with choice>
		seen := make(map[rune]int)
		for _, answers := range group {
			for _, choice := range answers {
				seen[choice]++
			}
		}

		// If the count of answers with choice matches the
		// total number of answers, that choice was included
		// in all answers.
		for k, v := range seen {
			if len(group) == v {
				counts[k]++
			}
		}
	}

	var groupCount int
	for _, v := range counts {
		groupCount += v
	}

	return groupCount
}

func parseGroupAnswers(input []string) [][]string {

	var groups [][]string
	var group []string

	for _, line := range input {

		// A line consistenting only of "-" delinates the separation between groups
		if line == "-" {
			groups = append(groups, group)
			group = []string{}
			continue
		}

		group = append(group, line)
	}

	// Add the final group, if one exists
	if len(group) > 0 {
		groups = append(groups, group)
	}

	return groups
}
