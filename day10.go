package main

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

type counter struct {
	value int
}

func main() {
	file, err := os.Open("input")
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
		//count := bfs(m.lights, m.buttons)
		count := optimal_parity(m.lights, m.buttons)
		total += count
	}

	fmt.Printf("Part 1 total = %d\n", total)

	// Part 2
	total = 0
	for _, m := range machines {
		total += apply_buttons_to_get_joltage(m.joltages, m.buttons)
	}
	fmt.Printf("Part 2 total = %d\n", total)
}

func bfs(target int, buttons []int) int {
	stack := make([]int, len(buttons))
	count := 0
	for {
		count++
		results := []int{}
		for _, v := range stack {
			for _, b := range buttons {
				result := v ^ b
				if result == target {
					return count
				}
				if !slices.Contains(results, result) {
					results = append(results, result)
				}
			}
		}
		stack = make([]int, len(results))
		copy(stack, results)
	}
}

func optimal_parity(target int, buttons []int) int {
	num_buttons := len(buttons)
	nmin := 1000000
	var loop func(v, bid, n int)
	loop = func(v, bid, n int) {
		if v == target {
			if n < nmin {
				nmin = n
			}
			return
		}
		if bid == num_buttons {
			return
		}
		// Apply button 0 or 1 times
		loop(v, bid+1, n)
		loop(v^buttons[bid], bid+1, n+1)
	}
	loop(0, 0, 0)
	return nmin
}

func apply_buttons_to_get_joltage(target []int, buttons []button) int {
	stack := [][]int{}
	for range len(buttons) {
		stack = append(stack, make([]int, len(target)))
	}
	count := 0
	for {
		count++
		results := [][]int{}
		for _, v := range stack {
			for _, b := range buttons {
				result := make([]int, len(target))
				copy(result, v)
				for _, bid := range b {
					result[bid]++
				}
				if slices.Equal(result, target) {
					return count
				}
				if !slices.ContainsFunc(results, func(s []int) bool {
					return slices.Equal(s, result)
				}) {
					results = append(results, result)
				}
			}
		}
		stack = make([][]int, len(results))
		copy(stack, results)
	}
}

func solve(joltages []int, buttons []button) int {
	// https://en.wikipedia.org/wiki/Gaussian_elimination
	// Make slice of coefficients represented as ints.
	cfs := make([]int, len(joltages))
	bit := 1
	for _, b := range buttons {
		for _, jid := range b {
			cfs[jid] += bit
		}
		bit *= 2
	}
	// Sort into row echelon form
	for i := range cfs {
		cfs[i] = cfs[i]*1000 + joltages[i]
	}
	slices.Sort(cfs)
	slices.Reverse(cfs)
	for i := range cfs {
		joltages[i] = cfs[i] % 1000
		cfs[i] /= 1000
	}

	// Perform row reduction
	for i := range len(cfs) - 1 {
		a := cfs[i]
		for j := i + 1; j < len(cfs); j++ {
			b := cfs[j]
			if a&b == b {
				cfs[i] -= b
				joltages[i] -= joltages[j]
			}
		}
	}

	// Find maximum number of allowable presses per button
	maxp := make([]int, len(buttons))
	bit = 1
	for i := range buttons {
		for j, cf := range cfs {
			if cf&bit > 0 {
				jv := joltages[j]
				if maxp[i] == 0 {
					maxp[i] = jv
				} else {
					if jv < maxp[i] {
						maxp[i] = jv
					}
				}
			}
		}
		bit *= 2
	}

	span := 1
	for _, mp := range maxp {
		span *= mp
	}
	return 3
}

func set_jolts(joltages []int, buttons []button) int {
	jolts := make([]int, len(joltages))
	min_count := 99999999
	// Find maximum number of allowable presses per button
	maxp := make([]int, len(buttons))
	for idx, b := range buttons {
		for i, j := range b {
			mp := joltages[j] + 1 // Add 1 since will use with % operator later
			if i == 0 {
				maxp[idx] = mp
			} else {
				if maxp[idx] > mp {
					maxp[idx] = mp
				}
			}
		}
	}
	span := 1
	for _, mp := range maxp {
		span *= mp
	}
	for n := range span {
		count := 0
		pn := n
	button_loop:
		for i, b := range buttons {
			presses := pn % maxp[i]
			pn /= maxp[i]
			count += presses
			for _, idx := range b {
				jolts[idx] += presses
				if jolts[idx] > joltages[idx] {
					break button_loop
				}
			}
		}
		if slices.Equal(joltages, jolts) {
			if count < min_count {
				min_count = count
			}
		}
		for i := range jolts {
			jolts[i] = 0
		}
	}
	return min_count
}
