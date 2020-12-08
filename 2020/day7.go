package main

import (
	"regexp"
	"strconv"
	"strings"
)

type bagRule struct {
	BagColor       string
	ContainingBags []innerBagRule
}

type innerBagRule struct {
	Count   int
	BagRule *bagRule
}

type bagRules map[string]*bagRule

var matchBagRule = regexp.MustCompile(`^([\w ]+) bags contain ([\w, ]+)\.$`)
var matchInnerBag = regexp.MustCompile(`([0-9]+) ([\w ]+) bag`)

func ContainingBags(input []string) int {

	bagRules := parseBagRules(input)

	var count int

	for _, bagRule := range bagRules {
		if recursiveBagSearch("shiny gold", bagRule) {
			count++
		}
	}

	return count
}

func BagsContainedInSingleBag(input []string) int {

	bagRules := parseBagRules(input)

	bagRule := bagRules["shiny gold"]
	return recursiveBagCount(bagRule)
}

func recursiveBagCount(bagRule *bagRule) int {

	var count int

	for _, containingBag := range bagRule.ContainingBags {
		count += containingBag.Count
		count += (containingBag.Count * recursiveBagCount(containingBag.BagRule))
	}

	return count
}

func recursiveBagSearch(target string, bagRule *bagRule) bool {

	for _, containingBag := range bagRule.ContainingBags {
		if containingBag.BagRule.BagColor == target {
			return true
		}

		if recursiveBagSearch(target, containingBag.BagRule) {
			return true
		}
	}

	return false
}

func parseBagRules(input []string) bagRules {

	bagRules := make(bagRules)

	for _, line := range input {

		matches := matchBagRule.FindStringSubmatch(line)
		bagColor := matches[1]
		outerRule := getOrCreateBagRule(bagRules, bagColor)

		innerRules := strings.Split(matches[2], ",")
		for _, innerRule := range innerRules {

			matches = matchInnerBag.FindStringSubmatch(innerRule)
			if matches == nil {
				// "no other bags"
				continue
			}

			innerBagCount, _ := strconv.Atoi(matches[1])
			innerBagColor := matches[2]
			innerRule := getOrCreateBagRule(bagRules, innerBagColor)

			outerRule.ContainingBags = append(outerRule.ContainingBags, innerBagRule{
				Count:   innerBagCount,
				BagRule: innerRule,
			})
		}
	}

	return bagRules
}

func getOrCreateBagRule(bagRules bagRules, bagColor string) *bagRule {

	var rule *bagRule

	if r, ok := bagRules[bagColor]; ok {
		rule = r
	} else {
		rule = &bagRule{
			BagColor: bagColor,
		}

		bagRules[bagColor] = rule
	}

	return rule
}
