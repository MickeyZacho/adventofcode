package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func dec1() {
	f, err := os.Open("dec-1-puzzle.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		fmt.Println(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lines []string // Array to store the lines
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line) // Append each line to the array
	}
	start_time := time.Now()
	calibration_sum := 0
	for _, line := range lines {
		line = strings.ReplaceAll(line, "twone", "21")
		line = strings.ReplaceAll(line, "oneight", "18")
		line = strings.ReplaceAll(line, "threeight", "38")
		line = strings.ReplaceAll(line, "fiveight", "58")
		line = strings.ReplaceAll(line, "nineight", "98")
		line = strings.ReplaceAll(line, "eightwo", "82")
		line = strings.ReplaceAll(line, "eighthree", "83")
		line = strings.ReplaceAll(line, "sevenine", "79")
		line = strings.ReplaceAll(line, "one", "1")
		line = strings.ReplaceAll(line, "two", "2")
		line = strings.ReplaceAll(line, "three", "3")
		line = strings.ReplaceAll(line, "four", "4")
		line = strings.ReplaceAll(line, "five", "5")
		line = strings.ReplaceAll(line, "six", "6")
		line = strings.ReplaceAll(line, "seven", "7")
		line = strings.ReplaceAll(line, "eight", "8")
		line = strings.ReplaceAll(line, "nine", "9")
		re := regexp.MustCompile(`[0-9]`)
		matches := re.FindAllString(line, -1)
		first_digit, _ := strconv.Atoi(matches[0])
		last_digit, _ := strconv.Atoi(matches[len(matches)-1])
		calibration_value := first_digit*10 + last_digit
		calibration_sum += calibration_value

		// fmt.Println("line: ", line, "\nfirst_digit: ", first_digit, "last_digit: ", last_digit, "calibration_value: ", calibration_value, "calibration_sum: ", calibration_sum)
		// fmt.Println("matches: ", matches)
	}
	fmt.Println(time.Now().Sub(start_time))
	fmt.Println(calibration_sum)
}
