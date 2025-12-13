package main

// This code works with the example data but part 2 takes too long to run with the input data

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type vertex struct {
	x, y int
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func RectArea(a, b vertex) int {
	return (AbsInt(b.x-a.x) + 1) * (AbsInt(b.y-a.y) + 1)
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var vertices []vertex
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		digits := strings.Split(line, ",")
		x, _ := strconv.Atoi(digits[0])
		y, _ := strconv.Atoi(digits[1])
		vertices = append(vertices, vertex{x, y})
	}

	// Part 1
	max_area := 0
	for i := range len(vertices) - 1 {
		for j := i + 1; j < len(vertices); j++ {
			a := RectArea(vertices[i], vertices[j])
			if a > max_area {
				max_area = a
			}
		}
	}
	fmt.Printf("Part 1 max area = %d\n", max_area)

	// Part 2

	// The vertices form a closed path containing the green tiles.
	// Rectangles formed by 2 vertices have 2 other corners that must be on or inside the closed path.

	vertices = append(vertices, vertices[0])

	// Loop finding the max area rectangle that has not already been rejected
	rejected_rects := []vertex{}
	for {
		max_area = 0
		var max_ids vertex
		var v1, v2 vertex
		for i := range len(vertices) - 1 {
			for j := i + 1; j < len(vertices); j++ {
				current_ids := vertex{i, j}
				if slices.Contains(rejected_rects, current_ids) {
					continue
				}
				a := RectArea(vertices[i], vertices[j])
				if a > max_area {
					max_area = a
					v1 = vertices[i]
					v2 = vertices[j]
					max_ids = current_ids
				}
			}
		}
		// Test points along the edges of the max rect
		ok := true
		// Top
		if v2.x > v1.x {
			for x := v1.x + 1; x <= v2.x; x++ {
				if !isInsidePolygon(vertices, vertex{x, v1.y}) {
					ok = false
					break
				}
			}
		} else {
			for x := v2.x; x < v1.x; x++ {
				if !isInsidePolygon(vertices, vertex{x, v1.y}) {
					ok = false
					break
				}
			}
		}
		if !ok {
			goto end
		}
		// Bottom
		if v2.x > v1.x {
			for x := v1.x; x < v2.x; x++ {
				if !isInsidePolygon(vertices, vertex{x, v2.y}) {
					ok = false
					break
				}
			}
		} else {
			for x := v2.x + 1; x <= v1.x; x++ {
				if !isInsidePolygon(vertices, vertex{x, v2.y}) {
					ok = false
					break
				}
			}
		}
		if !ok {
			goto end
		}
		// Left
		if v2.y > v1.y {
			for y := v1.y + 1; y <= v2.y; y++ {
				if !isInsidePolygon(vertices, vertex{v1.x, y}) {
					ok = false
					break
				}
			}
		} else {
			for y := v2.y; y < v1.y; y++ {
				if !isInsidePolygon(vertices, vertex{v1.x, y}) {
					ok = false
					break
				}
			}
		}
		if !ok {
			goto end
		}
		// Right
		if v2.y > v1.y {
			for y := v1.y; y < v2.y; y++ {
				if !isInsidePolygon(vertices, vertex{v2.x, y}) {
					ok = false
					break
				}
			}
		} else {
			for y := v2.y + 1; y <= v1.y; y++ {
				if !isInsidePolygon(vertices, vertex{v2.x, y}) {
					ok = false
					break
				}
			}
		}
	end:
		if ok {
			fmt.Println(v1, v2)
			fmt.Printf("Part 2 max area = %d\n", max_area)
			break
		} else {
			rejected_rects = append(rejected_rects, max_ids)
		}
	}
}

// Want to count how many times the point cuts though an edge vertically.
// If the point sits on a an edge then it is inside or
// if there is an odd number of traversals of edges up or down, then the point is inside.
func isInsidePolygon(verts []vertex, p vertex) bool {
	n := 0
	m := 0
	on_edge := false
	last_x := -1
	for i := range len(verts) - 1 {
		a := verts[i]
		b := verts[i+1]
		if a.y == b.y { // Horizontal line
			if a.x > b.x {
				a = b
				b = verts[i]
			}
			if a.x != last_x && p.x >= a.x && p.x <= b.x {
				if p.y < a.y {
					n++
				}
				if p.y == a.y {
					on_edge = true
					break
				}
				if p.y > a.y {
					m++
				}
			}
			last_x = b.x
		}
	}
	return on_edge || is_odd(n) || is_odd(m)
}

func is_odd(n int) bool {
	return n%2 != 0
}
