package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	EMPTY = iota
	ROLL_TO_GO
	ROLL
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

	num_rows := len(lines)
	num_cols := len(lines[0])

	rolls := make([][]int, num_rows)
	for i := 0; i < num_rows; i++ {
		rolls[i] = make([]int, num_cols)
	}
	for i := 0; i < num_rows; i++ {
		for j := 0; j < num_cols; j++ {
			if lines[i][j] == '@' {
				rolls[i][j] = ROLL
			}
		}
	}

	total := count_rolls(rolls, num_rows, num_cols)
	fmt.Printf("Part 1 total: %d\n", total)

	for {
		remove_rolls(rolls, num_rows, num_cols)
		num_rolls := count_rolls(rolls, num_rows, num_cols)
		if num_rolls == 0 {
			break
		}
		total += num_rolls
	}
	fmt.Printf("Part 2 total: %d\n", total)
}

func count_rolls(rolls [][]int, num_rows int, num_cols int) int {
	num_rolls := 0

	for i := 0; i < num_rows; i++ {
		for j := 0; j < num_cols; j++ {
			if rolls[i][j] == ROLL {
				num_adj := 0
				for row := i - 1; row <= i+1; row++ {
					if row >= 0 && row < num_rows {
						for col := j - 1; col <= j+1; col++ {
							if col >= 0 && col < num_cols {
								if row != i || col != j {
									if rolls[row][col] != EMPTY {
										num_adj++
									}
								}
							}
						}
					}
				}
				if num_adj < 4 {
					num_rolls++
					rolls[i][j] = ROLL_TO_GO
				}
			}
		}
	}
	return num_rolls
}

func remove_rolls(rolls [][]int, num_rows int, num_cols int) {
	for i := 0; i < num_rows; i++ {
		for j := 0; j < num_cols; j++ {
			if rolls[i][j] == ROLL_TO_GO {
				rolls[i][j] = EMPTY
			}
		}
	}
}
