package main

import (
	"strconv"
	"strings"
)

type replacementInstruction struct {
	programCounter int
	newInstruction string
}

func ExecuteProgram(instructions []string) int {

	var programCounter int
	var globalInt int

	executedLines := make(map[int]bool)

	for {
		if programCounter > len(instructions)-1 || executedLines[programCounter] {
			break
		}

		executedLines[programCounter] = true

		instruction := instructions[programCounter]
		operation, argument := parseInstruction(instruction)

		switch operation {
		case "jmp":
			programCounter += argument
		case "acc":
			globalInt += argument
			programCounter++
		case "nop":
			programCounter++
		}
	}

	return globalInt
}

func ExecuteProgramCleanly(instructions []string) int {

	var globalInt int
	var executedCleanly bool
	var programCounter int

	executedLines := make(map[int]bool)
	replacements := generateReplacementInstructions(instructions)

	// Replace jmp with nop and nop with jmp, one at a time, until the
	// program executes cleanly (does not result in an infinite loop)
	for _, replacement := range replacements {

		originalInstruction := instructions[replacement.programCounter]
		instructions[replacement.programCounter] = replacement.newInstruction

		// Execute program
		for {
			// We reached the end of the instructions: program executed cleanly!
			if programCounter > len(instructions)-1 {
				executedCleanly = true
				break
			}

			// Detected infinite loop: program did not execute cleanly
			if executedLines[programCounter] {
				break
			}

			executedLines[programCounter] = true

			instruction := instructions[programCounter]
			operation, argument := parseInstruction(instruction)

			// fmt.Println(programCounter, operation, argument)

			switch operation {
			case "jmp":
				programCounter += argument
			case "acc":
				globalInt += argument
				programCounter++
			case "nop":
				programCounter++
			}
		}

		if executedCleanly {
			break
		}

		// Reset variables to prepare to run the program again
		globalInt = 0
		programCounter = 0
		executedLines = map[int]bool{}
		instructions[replacement.programCounter] = originalInstruction
	}

	return globalInt
}

func parseInstruction(instruction string) (string, int) {

	parts := strings.Split(instruction, " ")
	operation := parts[0]
	argument, _ := strconv.Atoi(parts[1])

	return operation, argument
}

func generateReplacementInstructions(instructions []string) []replacementInstruction {

	var replacements []replacementInstruction

	for i, v := range instructions {
		operation, _ := parseInstruction(v)

		switch operation {
		case "jmp":
			replacements = append(replacements, replacementInstruction{
				programCounter: i,
				newInstruction: strings.Replace(v, "jmp", "nop", 1),
			})
		case "nop":
			replacements = append(replacements, replacementInstruction{
				programCounter: i,
				newInstruction: strings.Replace(v, "nop", "jmp", 1),
			})
		}
	}

	return replacements
}
