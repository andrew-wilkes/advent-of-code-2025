package main // This code is too inefficient for part 2 with my input data

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type button []int

type machine struct {
	lights   []bool
	buttons  []button
	joltages []int
}

// Binary encoded machine
type bmachine struct {
	lights   int
	buttons  []int
	joltages []int
}

func main() {
	file, err := os.Open("test")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var machines []machine
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		m := machine{}
		for _, chr := range parts[0] {
			switch chr {
			case '.':
				m.lights = append(m.lights, false)
			case '#':
				m.lights = append(m.lights, true)
			}
		}
		// Observe that switch digits appear to be single digits
		for i := 1; i < len(parts)-1; i++ {
			b := button{}
			for _, chr := range parts[i] {
				num, err := strconv.Atoi(string(chr))
				if err == nil {
					b = append(b, num)
				}
			}
			m.buttons = append(m.buttons, b)
		}
		jolts := strings.SplitSeq(strings.Trim(parts[len(parts)-1], "{}"), ",")
		for jolt := range jolts {
			jv, _ := strconv.Atoi(jolt)
			m.joltages = append(m.joltages, jv)
		}
		machines = append(machines, m)
	}

	// Binary encode the machines for ease of toggling values
	var bms []bmachine

	for _, m := range machines {
		bm := bmachine{}
		n := 1
		for _, d := range m.lights {
			if d {
				bm.lights += n
			}
			n *= 2
		}
		for _, b := range m.buttons {
			button := 0
			for _, n := range b {
				button += int(math.Pow(2, float64(n)))
			}
			bm.buttons = append(bm.buttons, button)
		}
		bms = append(bms, bm)
	}
	// Want the least number of button presses to set the lights correctly
	total := 0
	for _, m := range bms {
		count := apply_buttons(m.lights, m.buttons)
		total += count
	}

	fmt.Printf("Part 1 total = %d\n", total)

	// Part 2
	total = 0
	for _, m := range machines {
		total += set_jolts(m.joltages, m.buttons)
	}
	fmt.Printf("Part 2 total = %d\n", total)
}

func apply_buttons(target int, buttons []int) int {
	stack := make([]int, len(buttons))
	copy(stack, buttons) // Initial values after 1 press of the buttons
	count := 1
	if slices.Contains(stack, target) {
		return count
	}
	for {
		count++
		results := []int{}
		for _, v := range stack {
			for j := range len(buttons) {
				result := v ^ buttons[j]
				if result == target {
					return count
				}
				results = append(results, result)
			}
		}
		stack = make([]int, len(results))
		copy(stack, results)
	}
}

func set_jolts(joltages []int, buttons []button) int {
	// Sort buttons in decreasing order of impactfulness
	slices.SortFunc(buttons, func(a, b button) int {
		return len(b) - len(a)
	})
	stack := make([][]int, len(buttons))
	for i, b := range buttons {
		jolts := make([]int, len(joltages))
		for _, idx := range b {
			jolts[idx] = 1
		}
		stack[i] = jolts
	}
	count := 1
	if count > 10000 {
		return count
	}
	for {
		count++
		results := [][]int{}
		for _, jolts := range stack {
			for j := range len(buttons) {
				result := make([]int, len(joltages))
				copy(result, jolts)
				ok := true
				for _, idx := range buttons[j] {
					result[idx]++
					if result[idx] > joltages[idx] {
						ok = false
						break
					}
				}
				if ok {
					if slices.Equal(result, joltages) {
						return count
					}
					results = append(results, result)
				}
			}
		}
		stack = make([][]int, len(results))
		copy(stack, results)
	}
}
