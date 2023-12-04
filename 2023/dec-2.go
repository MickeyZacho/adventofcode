package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func dec2() {
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
	power_sum := 0
	for _, line := range lines {
		id_and_games := strings.Split(line, ":")
		id_and_trash := strings.Split(id_and_games[0], " ")
		id, _ := strconv.Atoi(id_and_trash[1])
		games := strings.Split(id_and_games[1], ";")
		line_possible := true
		red_val := 0
		green_val := 0
		blue_val := 0
		for _, game := range games {
			values := strings.Split(game, ",")
			for _, value := range values {
				color_val := strings.Split(value, " ")
				val, _ := strconv.Atoi(color_val[1])
				color := color_val[2]
				switch color {
				case "red":
					if val > 12 {
						line_possible = false
					}
					if val > red_val {
						red_val = val
					}
				case "green":
					if val > 13 {
						line_possible = false
					}
					if val > green_val {
						green_val = val
					}
				case "blue":
					if val > 14 {
						line_possible = false
					}
					if val > blue_val {
						blue_val = val
					}
				}
			}
		}
		if line_possible {
			sum += id
		}
		power := red_val * green_val * blue_val
		power_sum += power
	}
	fmt.Println(sum)
	fmt.Println(power_sum)
}
