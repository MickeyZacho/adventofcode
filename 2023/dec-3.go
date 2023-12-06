package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	f, err := os.Open("dec-3-puzzle.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		fmt.Println(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lines := make([]string, 3)
	index := 0
	for scanner.Scan() {
		line := scanner.Text()
		lines[index%3] = line
		index++
		first := (index % 3)
		second := (index + 1) % 3
		third := (index + 2) % 3

		fmt.Println("1: " + lines[first])
		fmt.Println("2: " + lines[second])
		fmt.Println("3: " + lines[third])

		regexp_symbols := regexp.MustCompile(`[^.0-9]`)
		regexp_numbers := regexp.MustCompile(`[0-9]+`)

		symbols := regexp_symbols.FindAllString(lines[second], -1)
		symbol_indexes := regexp_symbols.FindAllStringIndex(lines[second], -1)
		numbers := regexp_numbers.FindAllString(lines[second], -1)
		number_indexes := regexp_numbers.FindAllStringIndex(lines[second], -1)

		fmt.Println(symbols)
		fmt.Println(symbol_indexes)
		fmt.Println(numbers)
		fmt.Println(number_indexes)

		sum := 0

		for i, number_index := range number_indexes {

			start_index := number_index[0]
			end_index := number_index[1]
			for _, symbol_index := range symbol_indexes {

				if symbol_index[0] >= start_index && symbol_index[0] <= end_index || symbol_index[1] >= start_index && symbol_index[1] <= end_index {
					val, _ := strconv.Atoi(numbers[i])
					sum += val
				}
			}
		}

	}

}
