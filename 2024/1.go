package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func main() {
	dec1()
	dec1p2()
	dec2p1()
	dec2p2()
	dec3p1()
	dec3p2()
	dec4p1()
	dec4p2()
}

func dec1() {
	lines := get_input("1")
	left, right := process_input(lines)

	diff_sum := 0
	for i, l := range left {
		diff := int_abs(l - right[i])
		diff_sum += int(diff)
	}
	fmt.Println(diff_sum)
}

func get_input(number string) []string {
	f, err := os.Open(number + "-input.txt")
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

func process_input(lines []string) ([]int, []int) {
	left := []int{}
	right := []int{}
	for _, line := range lines {
		line_array := strings.Split(line, "   ")
		// fmt.Println(line_array[0])
		num, err := strconv.Atoi(line_array[0])
		if err != nil {
			fmt.Println("Error converting string to int:", err)
			continue
		}
		left = append(left, num)

		num, err = strconv.Atoi(line_array[1])
		if err != nil {
			fmt.Println("Error converting string to int:", err)
			continue
		}
		right = append(right, num)
	}

	slices.Sort(left)
	slices.Sort(right)

	return left, right
}

func dec1p2() {
	lines := get_input("1")
	left, right := process_input(lines)

	similarity_score := 0
	for _, l := range left {
		hits := 0
		for _, r := range right {
			if l == r {
				hits++
			}
		}
		similarity_score += hits * l
	}
	fmt.Println(similarity_score)

}

func dec2p1() {
	reports := get_input("2")
	valid_reports := 0
	for _, report := range reports {
		levels := strings.Split(report, " ")
		if checkReport(levels) {
			valid_reports++
		}
	}

	fmt.Println(valid_reports)
}

func dec2p2() {
	reports := get_input("2")
	valid_reports := 0
	for _, report := range reports {
		levels := strings.Split(report, " ")
		report_valid := false
		if checkReport(levels) {
			report_valid = true
		}
		for i, _ := range levels {
			// fmt.Println(levels)
			levels_minus_one := slices.Concat(levels[:i], levels[i+1:])
			if checkReport(levels_minus_one) {
				report_valid = true
			}
		}
		if report_valid {
			valid_reports++
		}
	}

	fmt.Println(valid_reports)
}

func checkReport(levels []string) bool {
	is_valid := true
	is_increasing := false
	is_decreasing := false

	for i, _ := range levels {
		if i == 0 {
			continue
		}
		a, _ := strconv.Atoi(levels[i-1])
		b, _ := strconv.Atoi(levels[i])
		if a > b {
			is_decreasing = true
			diff := a - b
			if diff > 3 {
				is_valid = false
			}
			if diff == 0 {
				is_valid = false
			}
		}
		if a < b {
			is_increasing = true
			diff := b - a
			if diff > 3 {
				is_valid = false
			}
			if diff == 0 {
				is_valid = false
			}
		}
		if a == b {
			is_valid = false
		}
	}
	if is_decreasing && is_increasing {
		is_valid = false
	}
	return is_valid
}

func int_abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func dec3p1() {
	lines := get_input("3")

	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)

	exprs := []string{}

	for _, line := range lines {
		exprs = append(re.FindAllString(line, -1), exprs...)
		// fmt.Println(exprs)
	}

	// fmt.Println(exprs)
	sum := 0
	for _, expr := range exprs {
		re2 := regexp.MustCompile(`\d{1,3}`)
		nums := re2.FindAllString(expr, -1)
		firstnum, _ := strconv.Atoi(nums[0])
		secondnum, _ := strconv.Atoi(nums[1])
		sum += (firstnum * secondnum)
	}
	fmt.Println(sum)
}

func dec3p2() {
	lines := get_input("3")

	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`)

	exprs := []string{}

	for _, line := range lines {
		exprs = append(exprs, re.FindAllString(line, -1)...)
		// fmt.Println(exprs)
	}

	// fmt.Println(exprs)
	sum := 0
	enabled := true
	for _, expr := range exprs {
		if expr == "do()" {
			// fmt.Println("do")
			enabled = true
			continue
		}
		if expr == "don't()" {
			// fmt.Println("don't")
			enabled = false
			continue
		}
		if !enabled {
			continue
		}
		// fmt.Println(expr)
		re2 := regexp.MustCompile(`\d{1,3}`)
		nums := re2.FindAllString(expr, -1)
		firstnum, _ := strconv.Atoi(nums[0])
		secondnum, _ := strconv.Atoi(nums[1])
		sum += (firstnum * secondnum)
	}
	fmt.Println(sum)
}

func dec4p1() {
	lines := get_input("4")
	// fmt.Println(lines)
	count := 0
	xcount := 0
	for i, line := range lines {
		// fmt.Println(line)
		n := len(line)
		m := len(lines)
		// fmt.Println(m)
		// fmt.Println(n)
		for j := 0; j < n-1; j++ {
			// fmt.Println(string(line[j]))
			// fmt.Println(j)
			// fmt.Println(i)
			if string(line[j]) == "X" {
				xcount++
				lookforwards := j < n-3
				lookbackwards := j >= 3
				lookupwards := i >= 3
				lookdownwards := i < m-3

				if lookforwards && string(line[j+1]) == "M" && string(line[j+2]) == "A" && string(line[j+3]) == "S" {
					count++
				}
				if lookbackwards && string(line[j-1]) == "M" && string(line[j-2]) == "A" && string(line[j-3]) == "S" {
					count++
				}
				if lookupwards && string(lines[i-1][j]) == "M" && string(lines[i-2][j]) == "A" && string(lines[i-3][j]) == "S" {
					count++
				}
				if lookdownwards && string(lines[i+1][j]) == "M" && string(lines[i+2][j]) == "A" && string(lines[i+3][j]) == "S" {
					count++
				}
				if lookdownwards && lookforwards && string(lines[i+1][j+1]) == "M" && string(lines[i+2][j+2]) == "A" && string(lines[i+3][j+3]) == "S" {
					count++
				}
				if lookupwards && lookbackwards && string(lines[i-1][j-1]) == "M" && string(lines[i-2][j-2]) == "A" && string(lines[i-3][j-3]) == "S" {
					count++
				}
				if lookdownwards && lookbackwards && string(lines[i+1][j-1]) == "M" && string(lines[i+2][j-2]) == "A" && string(lines[i+3][j-3]) == "S" {
					count++
				}
				if lookupwards && lookforwards && string(lines[i-1][j+1]) == "M" && string(lines[i-2][j+2]) == "A" && string(lines[i-3][j+3]) == "S" {
					// fmt.Println("Found MAS")
					count++
				}
			}
		}
	}
	fmt.Println(count)
	// fmt.Println(xcount)
}

func dec4p2() {

}
