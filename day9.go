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

type rect struct {
	a, b vertex
	area int
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

	rects := []rect{}
	// Part 1
	max_area := 0
	for i := range len(vertices) - 1 {
		for j := i + 1; j < len(vertices); j++ {
			a := RectArea(vertices[i], vertices[j])
			rects = append(rects, get_rect(vertices[i], vertices[j], a))
			if a > max_area {
				max_area = a
			}
		}
	}
	fmt.Printf("Part 1 max area = %d\n", max_area)

	// Part 2

	// Sort rectangles by area
	slices.SortFunc(rects, func(x, y rect) int {
		return y.area - x.area
	})

	// The vertices form a closed path containing the green tiles.
	vertices = append(vertices, vertices[0])

	max_area = 0
	start := len(rects) - 80000 // Skip some of the largest rects to save time
	if start < 0 {
		start = 0
	}
	for i := start; i < len(rects); i++ {
		rect := rects[i]
		if is_inside(rect, vertices) {
			max_area = rect.area
			break
		}
	}
	fmt.Printf("Part 2 max area = %d\n", max_area)
}

func get_rect(v1, v2 vertex, area int) rect {
	var a, b vertex
	if v1.x < v2.x {
		a.x = v1.x
		b.x = v2.x
	} else {
		a.x = v2.x
		b.x = v1.x
	}
	if v1.y < v2.y {
		a.y = v1.y
		b.y = v2.y
	} else {
		a.y = v2.y
		b.y = v1.y
	}
	return rect{a, b, area}
}

func is_inside(r rect, vertices []vertex) bool {
	// Horizontal edges
	for x := r.a.x; x <= r.b.x; x++ {
		if !point_inside(vertex{x, r.a.y}, vertices) {
			return false
		}
		if !point_inside(vertex{x, r.b.y}, vertices) {
			return false
		}
	}
	// Vertical edges
	for y := r.a.y + 1; y < r.b.y; y++ {
		if !point_inside(vertex{r.a.x, y}, vertices) {
			return false
		}
		if !point_inside(vertex{r.b.x, y}, vertices) {
			return false
		}
	}
	return true
}

func point_inside(p vertex, vertices []vertex) bool {
	odd := false
	for i := range len(vertices) - 1 {
		a := vertices[i]
		b := vertices[i+1]
		// Check if point is on the line
		if a.x == b.x {
			// Vertical line
			if p.x == a.x {
				if a.y < b.y {
					if p.y >= a.y && p.y <= b.y {
						return true
					}
				} else {
					if p.y >= b.y && p.y <= a.y {
						return true
					}
				}
			} else { // Check for crossing nodes to the left of the point
				if a.y < b.y {
					if a.x < p.x {
						if a.y < p.y && b.y >= p.y {
							odd = !odd
						}
					}
				} else {
					if a.x < p.x {
						if b.y < p.y && a.y >= p.y {
							odd = !odd
						}
					}
				}
			}
		} else {
			// Horizontal line
			if p.y == a.y {
				if a.x < b.x {
					if p.x >= a.x && p.x <= b.x {
						return true
					}
				} else {
					if p.x >= b.x && p.x <= a.x {
						return true
					}
				}
			}
		}
	}
	return odd
}
