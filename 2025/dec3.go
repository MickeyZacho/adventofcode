package main

import (
	"fmt"
	"os"
	"strconv"
)

func dec3() {
	banks := readInputFromFile("3")
	totalOutputJoltage := 0
	for _, bankLine := range banks {
		firstDigitContenders := bankLine[0 : len(bankLine)-1]
		firstDigit := 0
		index := 0
		for i := 0; i < len(firstDigitContenders); i++ {
			contender, err := strconv.Atoi(firstDigitContenders[i : i+1])
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			if contender > firstDigit {
				firstDigit = contender
				index = i
			}
		}
		secondDigitContenders := bankLine[index+1:]
		secondDigit := 0
		for j := 0; j < len(secondDigitContenders); j++ {
			contender, err := strconv.Atoi(secondDigitContenders[j : j+1])
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			if contender > secondDigit {
				secondDigit = contender
			}
		}
		joltageStr := strconv.Itoa(firstDigit) + strconv.Itoa(secondDigit)
		fmt.Println("Result for bank line " + bankLine + ": " + joltageStr)
		joltage := firstDigit*10 + secondDigit
		fmt.Println("Joltage: " + strconv.Itoa(joltage))

		totalOutputJoltage += joltage
	}
	fmt.Println("Total output joltage: " + strconv.Itoa(totalOutputJoltage))
}

func dec3_2() {
	banks := readInputFromFile("3")
	totalOutputJoltage := 0
	BATTERIES_PR_BANK := 12
	for _, bankLine := range banks {
		fmt.Println("Bank line: " + bankLine)
		digits := []int{}
		joltage := 0
		lastChosenIndex := -1
		for i := range BATTERIES_PR_BANK {
			fmt.Println("Choosing digit " + strconv.Itoa(i))
			fmt.Println("Last chosen index: " + strconv.Itoa(lastChosenIndex))
			nextDigitContenders := bankLine[lastChosenIndex+1 : len(bankLine)-(BATTERIES_PR_BANK-i-1)]
			fmt.Println("Next digit contenders: " + nextDigitContenders)
			fmt.Println("Length of next digit contenders: " + strconv.Itoa(len(nextDigitContenders)))
			nextDigit := 0
			nextDigitIndex := -1
			for j := 0; j < len(nextDigitContenders); j++ {
				indexContender := lastChosenIndex + 1 + j
				fmt.Println("next: " + nextDigitContenders[j:j+1])
				contender, err := strconv.Atoi(nextDigitContenders[j : j+1])
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				if contender > nextDigit {
					fmt.Println("Last chosen index: " + strconv.Itoa(lastChosenIndex))
					fmt.Println("Index contender: " + strconv.Itoa(indexContender))
					fmt.Println("Contender: " + strconv.Itoa(contender))
					fmt.Println("j: " + strconv.Itoa(j))
					nextDigit = contender
					nextDigitIndex = indexContender
				}
			}
			lastChosenIndex = nextDigitIndex
			fmt.Println("Chosen next digit: " + strconv.Itoa(nextDigit))
			fmt.Println("Chosen next digit index: " + strconv.Itoa(nextDigitIndex))

			digits = append(digits, nextDigit)

		}
		for k := 0; k < len(digits); k++ {
			joltage = joltage*10 + digits[k]
		}
		totalOutputJoltage += joltage
	}
	fmt.Println("Total output joltage: " + strconv.Itoa(totalOutputJoltage))
}
