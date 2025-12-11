package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
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
	start_pos := strings.IndexByte(lines[0], 'S')
	timeline_count := 0
	beams := map[int]int{}
	beams[start_pos] = 1
	num_rows := len(lines)
	num_splits := 0
	row := 2
	for row < num_rows {
		splitter_positions := find_splitters(lines[row])
		new_beams := map[int]int{}
		for beam_pos, weight := range beams {
			splitter_idx := slices.Index(splitter_positions, beam_pos)
			if splitter_idx > -1 {
				num_splits++
				new_beams[beam_pos-1] = new_beams[beam_pos-1] + weight // If key doesn't exist value = 0
				new_beams[beam_pos+1] = new_beams[beam_pos+1] + weight
			} else {
				new_beams[beam_pos] = new_beams[beam_pos] + weight
			}
		}
		beams = new_beams
		row += 2
	}

	fmt.Printf("Part 1. Number of splits = %d\n", num_splits)

	for _, weight := range beams {
		timeline_count += weight
	}

	fmt.Printf("Part 2. Number of timelines = %d\n", timeline_count)
}

func find_splitters(str string) []int {
	positions := []int{}
	for idx, chr := range str {
		if chr == '^' {
			positions = append(positions, idx)
		}
	}
	return positions
}
