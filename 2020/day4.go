package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var requiredFields = []string{
	"byr", // Birth Year
	"iyr", // Issue Year
	"eyr", // Expiration Year
	"hgt", // Height
	"hcl", // Hair Color
	"ecl", // Eye Color
	"pid", // Passport ID
	// "cid", // Country ID
}

func ValidPassports(input string) int {

	passports := parsePassportInput(input)

	var validPassports int
	for _, passport := range passports {
		if validatePassport(passport) {
			validPassports++
		}
	}

	return validPassports
}

func parsePassportInput(input string) []string {

	var passports []string
	var passport string

	// Break input into lines and extract the passports
	for _, line := range strings.Split(strings.TrimSuffix(input, "\n"), "\n") {

		if line == "" && passport != "" {
			passports = append(passports, passport)
			passport = ""
			continue
		}

		passport += line + " "
	}

	// Add the final passport, if one exists
	if passport != "" {
		passports = append(passports, passport)
	}

	return passports
}

func validatePassport(passport string) bool {

	for _, field := range requiredFields {

		// Invalid: the field is missing from this passport entirely
		if !strings.Contains(passport, field) {
			return false
		}

		var extractFieldValue = regexp.MustCompile(fmt.Sprintf(`%s:([^\s]+)`, field))
		matches := extractFieldValue.FindStringSubmatch(passport)

		// Invalid: formating error
		if matches == nil || len(matches) < 2 {
			return false
		}

		value := matches[1]

		var valid bool

		switch field {

		case "byr":
			valid = validateValueInt(value, 1920, 2002)

		case "iyr":
			valid = validateValueInt(value, 2010, 2020)

		case "eyr":
			valid = validateValueInt(value, 2020, 2030)

		case "hgt":
			hgtCmRegex := regexp.MustCompile(`^([0-9]+)cm$`)
			cmValid := hgtCmRegex.FindStringSubmatch(value)

			hgtInRegex := regexp.MustCompile(`^([0-9]+)in$`)
			inValid := hgtInRegex.FindStringSubmatch(value)

			if cmValid != nil {
				valid = validateValueInt(cmValid[1], 150, 193)
			} else if inValid != nil {
				valid = validateValueInt(inValid[1], 59, 76)
			}

		case "hcl":
			hclRegex := regexp.MustCompile(`^#[0-9a-f]{6}$`)
			valid = hclRegex.MatchString(value)

		case "ecl":
			if value == "amb" || value == "blu" || value == "brn" || value == "gry" || value == "grn" || value == "hzl" || value == "oth" {
				valid = true
			}

		case "pid":
			pidRegex := regexp.MustCompile(`^[0-9]{9}$`)
			valid = pidRegex.MatchString(value)
		}

		// Invalid: data validation failed for this field
		if !valid {
			return false
		}
	}

	return true
}

// validateValueInt validates that the given value converted to an int
// is within the min and max boundaries.
func validateValueInt(value string, min, max int) bool {
	i, _ := strconv.Atoi(value)
	if i >= min && i <= max {
		return true
	}

	return false
}
