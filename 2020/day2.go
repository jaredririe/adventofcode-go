package main

import (
	"regexp"
	"strconv"
	"strings"
)

var matchPasswordString = regexp.MustCompile("([0-9]+)-([0-9]+) ([a-z]+): ([a-z]+)")

func ValidPasswords(passwords []string) int {

	var valid int

	for _, password := range passwords {

		matches := matchPasswordString.FindStringSubmatch(password)
		minStr, maxStr, letter, pwd := matches[1], matches[2], matches[3], matches[4]

		min, _ := strconv.Atoi(minStr)
		max, _ := strconv.Atoi(maxStr)

		count := strings.Count(pwd, letter)
		if count >= min && count <= max {
			valid++
		}
	}

	return valid
}

func ValidPasswords2(passwords []string) int {

	var valid int

	for _, password := range passwords {

		matches := matchPasswordString.FindStringSubmatch(password)
		first, second, letter, pwd := matches[1], matches[2], matches[3], matches[4]

		firstIdx, _ := strconv.Atoi(first)
		secondIdx, _ := strconv.Atoi(second)

		firstIdx--
		secondIdx--

		firstLetter := string([]rune(pwd)[firstIdx])
		secondLetter := string([]rune(pwd)[secondIdx])

		// Valid if exactly one of the first letter or second letter matches
		m1 := firstLetter == letter && secondLetter != letter
		m2 := firstLetter != letter && secondLetter == letter
		if m1 || m2 {
			valid++
		}
	}

	return valid
}
