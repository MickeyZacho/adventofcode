package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input_lines := get_input("1-1")
	current_val := 50
	zeroes := 0
	zero_cross := 0
	for _, code := range input_lines {
		direction := code[0:1]
		steps, err := strconv.Atoi(code[1:])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		initial_val := current_val

		shortened_steps := steps % 100
		step_rotations := steps / 100
		zero_cross += step_rotations

		switch direction {
		case "R":
			current_val = (current_val + shortened_steps)
			if current_val > 99 {
				zero_cross++
				current_val -= 100
			}
		case "L":
			current_val = (current_val - shortened_steps)
			if current_val == 0 && initial_val != 0 {
				zero_cross++
			}
			if current_val < 0 && initial_val != 0 {
				zero_cross++
			}
			if current_val < 0 {
				current_val += 100
			}
		default:
			os.Exit(2)
		}

		if current_val == 0 {
			zeroes++
		}
	}
	fmt.Println("Result: " + strconv.Itoa(current_val))
	fmt.Println("task 1: " + strconv.Itoa(zeroes))
	fmt.Println("task 2: " + strconv.Itoa(zero_cross))
}

func get_input(number string) []string {
	f, err := os.Open("dec" + number + ".txt")
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
	return lines
}
