package main

import (
	"fmt"
	"os"
	"strconv"
)

func dec1() {
	inputLines := readInputFromFile("1-1")
	currentVal := 50
	zeroes := 0
	zeroCross := 0
	for _, code := range inputLines {
		direction := code[0:1]
		steps, err := strconv.Atoi(code[1:])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		initialVal := currentVal

		shortenedSteps := steps % 100
		stepRotations := steps / 100
		zeroCross += stepRotations

		switch direction {
		case "R":
			currentVal = (currentVal + shortenedSteps)
			if currentVal > 99 {
				zeroCross++
				currentVal -= 100
			}
		case "L":
			currentVal = (currentVal - shortenedSteps)
			if currentVal == 0 && initialVal != 0 {
				zeroCross++
			}
			if currentVal < 0 && initialVal != 0 {
				zeroCross++
			}
			if currentVal < 0 {
				currentVal += 100
			}
		default:
			os.Exit(2)
		}

		if currentVal == 0 {
			zeroes++
		}
	}
	fmt.Println("Result: " + strconv.Itoa(currentVal))
	fmt.Println("task 1: " + strconv.Itoa(zeroes))
	fmt.Println("task 2: " + strconv.Itoa(zeroCross))
}
