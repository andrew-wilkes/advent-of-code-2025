package main

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

func TriArea(a, b, c vertex) int {
	// Not /2 since only interested in the sign
	return (b.x-a.x)*(c.y-a.y) - (c.x-a.x)*(b.y-a.y)
}

func main() {
	file, err := os.Open("test")
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
			goto end // Goto :)
		}
		// Bottom
		if v2.x > v1.x {
			for x := v1.x + 1; x <= v2.x; x++ {
				if !isInsidePolygon(vertices, vertex{x, v2.y}) {
					ok = false
					break
				}
			}
		} else {
			for x := v2.x; x < v1.x; x++ {
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
			for y := v1.y + 1; y <= v2.y; y++ {
				if !isInsidePolygon(vertices, vertex{v2.x, y}) {
					ok = false
					break
				}
			}
		} else {
			for y := v2.y; y < v1.y; y++ {
				if !isInsidePolygon(vertices, vertex{v2.x, y}) {
					ok = false
					break
				}
			}
		}
	end:
		if ok {
			fmt.Printf("Part2 max area = %d\n", max_area)
			break
		} else {
			rejected_rects = append(rejected_rects, max_ids)
		}
	}
}

// Want to count how many time the point cuts though an edge vertically.
// But if it sits on a top edge then get the wrong result.
func isInsidePolygon(verts []vertex, p vertex) bool {
	n := 0
	for i := range len(verts) - 1 {
		a := verts[i]
		b := verts[i+1]
		if a.y == b.y { // Horizontal line
			if a.x > b.x {
				a = b
				b = verts[i]
			}
			if p.y < a.y && p.x >= a.x && p.x <= b.x {
				// Line is below p
				n++
			}
		}
	}
	return n%2 != 0
}
