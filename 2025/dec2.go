package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func dec2() {
	// Input consists of a single line of ranges like "5-10,15-20"
	inputLines := get_input("2")
	rangesString := inputLines[0]
	ranges := strings.Split(rangesString, ",")
	invalidIdSum1 := 0
	invalidIdSum2 := 0
	for _, r := range ranges {
		rSplit := strings.Split(r, "-")
		rStart, err := strconv.Atoi(rSplit[0])
		if err != nil {
			fmt.Println("Err 1")
			os.Exit(1)
		}
		rEnd, err := strconv.Atoi(rSplit[1])
		if err != nil {
			fmt.Println("Err 2")
			os.Exit(1)
		}

		for i := rStart; i <= rEnd; i++ {
			if isInvalidInputDec2_1(i) {
				invalidIdSum1 += i
			}
			if isInvalidInputDec2_2(i) {
				invalidIdSum2 += i
			}
		}
	}

	fmt.Println("Dec2_1 result: " + strconv.Itoa(invalidIdSum1))
	fmt.Println("Dec2_2 result: " + strconv.Itoa(invalidIdSum2))

}

// Input is invalid if it is some repeating digits: i.e. "123123" or "6767"
func isInvalidInputDec2_1(i int) bool {
	indexString := strconv.Itoa(i)
	digits := len(indexString)

	// Early escape if input is uneven amount of digits
	if digits%2 == 1 {
		return false
	}

	firstDigits := indexString[0 : digits/2]
	lastDigits := indexString[digits/2 : digits]

	return firstDigits == lastDigits
}

// Input is invalid if it is a sequence of digits repeated at least twice
func isInvalidInputDec2_2(i int) bool {
	numStr := strconv.Itoa(i)
	digits := len(numStr)

	var sequenceLengths []int
	// find all numbers digits is divisible by
	for i := 1; i <= digits/2; i++ {
		if digits%i == 0 {
			sequenceLengths = append(sequenceLengths, i)
		}
	}

	// seqLen is the number of digits we are checking if is repeated
	for _, seqLen := range sequenceLengths {
		numRepeats := digits / seqLen // digits is divisble by seqLen
		sequence := numStr[0:seqLen]
		isSequenceValid := true
		for i := 1; i < numRepeats; i++ {
			if numStr[i*seqLen:(i+1)*seqLen] != sequence {
				isSequenceValid = false
			}
		}
		if isSequenceValid {
			return true
		}
	}

	return false
}
