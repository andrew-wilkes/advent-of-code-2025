package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Part 1
	total_jolts := 0
	for _, str := range lines {
		total_jolts += get_jolts(str)
	}
	fmt.Printf("Part 1. Total jolts = %d\n", total_jolts)
	// End of Part 1

	// Part 2
	total_jolts = 0
	for _, str := range lines {
		total_jolts += get_big_jolts(str)
	}
	fmt.Printf("Part 2. Total jolts = %d\n", total_jolts)
}

func get_jolts(bank string) int {
	j1 := 0
	j2 := 0
	// Find the largest value before the last value
	// Then find the largest value after the above value
	pos := 0
	for i := 0; i < len(bank)-1; i++ {
		num, _ := strconv.Atoi(string(bank[i]))
		if num > j1 {
			j1 = num
			pos = i
		}
	}
	for i := pos + 1; i < len(bank); i++ {
		num, _ := strconv.Atoi(string(bank[i]))
		if num > j2 {
			j2 = num
		}
	}
	return j1*10 + j2
}

func get_big_jolts(bank string) int {
	jolts := 0
	pos := -1
	togo := 11
	for i := 0; i < 12; i++ {
		biggest := 0
		for j := pos + 1; j < len(bank)-togo; j++ {
			num, _ := strconv.Atoi(string(bank[j]))
			if num > biggest {
				biggest = num
				pos = j
			}
		}
		togo--
		jolts = jolts*10 + biggest
	}
	return jolts
}
