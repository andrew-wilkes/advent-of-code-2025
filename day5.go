package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type num_pair struct {
	min, max int
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var ranges []num_pair
	num_fresh := 0
	getting_ranges := true
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if getting_ranges {
			if len(line) == 0 {
				getting_ranges = false
			} else {
				nums := strings.Split(line, "-")
				start_id, _ := strconv.Atoi(nums[0])
				end_id, _ := strconv.Atoi(nums[1])
				ranges = append(ranges, num_pair{start_id, end_id})
			}
		} else {
			num, _ := strconv.Atoi(line)
			if is_num_in_range(ranges, num) {
				num_fresh++
			}
		}
	}
	fmt.Printf("Part 1. Number of fresh items = %d\n", num_fresh)

	merged_ranges := merge_ranges(ranges)
	total := count_ids(merged_ranges)
	fmt.Printf("Part 2. Number of fresh ids = %d\n", total)
}

func is_num_in_range(ranges []num_pair, num int) bool {
	for _, np := range ranges {
		if num >= np.min && num <= np.max {
			return true
		}
	}
	return false
}

func merge_ranges(ranges []num_pair) []num_pair {
	// Sort the ranges in ascending order of min value
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].min < ranges[j].min
	})
	merged_intervals := []num_pair{}
	last_merged_idx := 0
	for _, current_interval := range ranges {
		if len(merged_intervals) == 0 {
			merged_intervals = append(merged_intervals, current_interval)
		} else {
			// Check for overlap
			if merged_intervals[last_merged_idx].max >= current_interval.min {
				if current_interval.max > merged_intervals[last_merged_idx].max {
					merged_intervals[last_merged_idx].max = current_interval.max
				}
			} else {
				merged_intervals = append(merged_intervals, current_interval)
				last_merged_idx++
			}
		}
	}
	return merged_intervals
}

func count_ids(ranges []num_pair) int {
	sum := 0
	for _, num_pair := range ranges {
		sum += num_pair.max - num_pair.min + 1
	}
	return sum
}
