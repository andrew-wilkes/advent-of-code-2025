package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type vec3 struct {
	x, y, z, id int
}

type pair struct {
	v1, v2, dist int
}

func dsquared(v1, v2 vec3) int {
	x := v1.x - v2.x
	y := v1.y - v2.y
	z := v1.z - v2.z
	return x*x + y*y + z*z
}

func main() {
	file, err := os.Open("test")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var junction_boxes []vec3
	scanner := bufio.NewScanner(file)
	id := 0
	for scanner.Scan() {
		line := scanner.Text()
		digits := strings.Split(line, ",")
		x, _ := strconv.Atoi(digits[0])
		y, _ := strconv.Atoi(digits[1])
		z, _ := strconv.Atoi(digits[2])
		junction_boxes = append(junction_boxes, vec3{x, y, z, id})
		id++
	}

	// Part 1
	num_pairs := 10
	var circuits [][]int
	num_junction_boxes := len(junction_boxes)

	// Get distances between each pair of vertices
	var pairs []pair
	for i := 0; i < num_junction_boxes-1; i++ {
		v1 := junction_boxes[i]
		for j := i + 1; j < num_junction_boxes; j++ {
			v2 := junction_boxes[j]
			pairs = append(pairs, pair{v1: i, v2: j, dist: dsquared(v1, v2)})
		}
	}

	// Sort into order of distance
	slices.SortFunc(pairs, func(a, b pair) int {
		return a.dist - b.dist
	})

	for n := range num_pairs {
		pair := pairs[n]
		create_circuit := true
		for idx, circ := range circuits {
			if slices.Contains(circ, pair.v1) {
				if !slices.Contains(circ, pair.v2) {
					// If any of the following circuits contain v2 append it
					if !append_circuit(circuits, idx, pair.v2) {
						circuits[idx] = append(circuits[idx], pair.v2)
					}
				}
				create_circuit = false
				break
			}
			if slices.Contains(circ, pair.v2) {
				if !slices.Contains(circ, pair.v1) {
					// If any of the following circuits contain v1 append it
					if !append_circuit(circuits, idx, pair.v1) {
						circuits[idx] = append(circuits[idx], pair.v1)
					}
				}
				create_circuit = false
				break
			}
		}
		if create_circuit {
			circuits = append(circuits, []int{pair.v1, pair.v2})
		}
	}
	// Multiply the three largest circuit sizes
	slices.SortFunc(circuits, func(a, b []int) int {
		return len(b) - len(a)
	})
	total := 1
	for i := range 3 {
		total *= len(circuits[i])
	}
	fmt.Printf("Part 1 total = %d\n", total)
}

func append_circuit(circuits [][]int, idx, v int) bool {
	appended := false
	for i := idx + 1; i < len(circuits); i++ {
		if slices.Contains(circuits[i], v) {
			circuits[idx] = append(circuits[idx], circuits[i]...)
			circuits[i] = []int{}
			appended = true
			break
		}
	}
	return appended
}
