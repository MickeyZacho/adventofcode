package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	f, err := os.Open("dec-2-puzzle.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		fmt.Println(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	sum := 0
	re := regexp.MustCompile(`[0-9]+`)
	for _, line := range lines {
		// match_indexes := re.FindAllStringIndex(line, -1)
		matches := re.FindAllString(line, -1)
		fmt.Println(matches)
	}
	fmt.Println(sum)

}
