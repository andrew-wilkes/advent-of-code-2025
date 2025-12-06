package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type num_aligned struct {
	n             int
	txt           string
	right_aligned bool
}

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

	// Determine column widths and set up data in columns
	var col_widths []int
	var cols [][]num_aligned
	var ops []byte
	for row, line := range lines {
		vals := strings.Fields(line)
		if len(col_widths) == 0 {
			col_widths = make([]int, len(vals))
			cols = make([][]num_aligned, len(vals))
			ops = make([]byte, len(vals))
		}
		for col, val := range vals {
			if row == 0 {
				cols[col] = []num_aligned{}
			}
			num, err := strconv.Atoi(val)
			if err != nil {
				ops[col] = val[0]
			} else {
				cols[col] = append(cols[col], num_aligned{n: num})
				width := len(val)
				if width > col_widths[col] {
					col_widths[col] = width
				}
			}
		}
	}

	// Set alignment properties.
	for row, line := range lines {
		if row == len(cols[0]) {
			break
		}
		idx := 0
		// Try to detect right alignment where a space prefixes the digits
		for col, width := range col_widths {
			sub_str := line[idx : idx+width]
			cols[col][row].txt = sub_str // Need this for part 2.
			if strings.HasPrefix(sub_str, " ") {
				for idx := range cols[col] {
					cols[col][idx].right_aligned = true // Didn't actually need this data.
				}
			}
			idx += width + 1
		}
	}

	// Part 1.
	total := 0
	for col, op := range ops {
		result := 0
		switch op {
		case '+':
			for _, num := range cols[col] {
				result += num.n
			}
		case '*':
			result = 1
			for _, num := range cols[col] {
				result *= num.n
			}
		}
		total += result
	}
	fmt.Printf("Part 1 total = %d\n", total)

	// Part 2.
	total = 0
	for col, width := range col_widths {
		nums := []int{}
		idx := width
		for idx > 0 {
			idx--
			cstr := ""
			for _, num := range cols[col] {
				chr := num.txt[idx : idx+1]
				if chr != " " {
					cstr = cstr + chr
				}
			}
			n, _ := strconv.Atoi(cstr)
			nums = append(nums, n)
		}
		result := 0
		switch ops[col] {
		case '+':
			for _, num := range nums {
				result += num
			}
		case '*':
			result = 1
			for _, num := range nums {
				result *= num
			}
		}
		total += result
	}
	fmt.Printf("Part 2 total = %d\n", total)
}
