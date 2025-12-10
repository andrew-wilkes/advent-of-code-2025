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

	// Part 2 - WIP

	// The vertices form a closed path containing the green tiles.
	// Rectangles formed by 2 vertices have 2 other corners that must be on or inside the closed path.

	// Vertices should be in counterclockwise order for use of my function to test isInsidePolygon
	slices.Reverse(vertices)

	vertices = append(vertices, vertices[0])

	rejected_rects := []vertex{}
	for {
		max_area := 0
		var max_ids vertex
		var v1, v2 vertex
		// Get the max area up to the ceiling
		for i := range len(vertices) - 1 {
			for j := i + 1; j < len(vertices); j++ {
				max_ids = vertex{i, j}
				if slices.Contains(rejected_rects, max_ids) {
					continue
				}
				a := RectArea(vertices[i], vertices[j])
				if a > max_area {
					max_area = a
					v1 = vertices[i]
					v2 = vertices[j]
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
			continue
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
			continue
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
			continue
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
		if ok {
			fmt.Printf("Part2 max area = %d\n", max_area)
			break
		} else {
			rejected_rects = append(rejected_rects, max_ids)
		}
	}
}

func isInsidePolygon(verts []vertex, p vertex) bool {
	for i := range len(verts) - 1 {
		if TriArea(p, verts[i], verts[i+1]) < 0 {
			return false
		}
	}
	return true
}
