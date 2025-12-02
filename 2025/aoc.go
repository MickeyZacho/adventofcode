package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	dec1()
	dec2()
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
