package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type region struct {
	width, length int
	presents      []int
}

type shape [][][3]int // A slice of variants of the base shape

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	shapes := []shape{}
	regions := []region{}
	areas := []int{}
	scanner := bufio.NewScanner(file)
	line_num := 0
	for scanner.Scan() {
		line := scanner.Text()
		block := line_num / 5
		item := line_num % 5
		if block < 6 {
			if item == 0 {
				shapes = append(shapes, shape{[][3]int{}})
				areas = append(areas, 0)
			} else {
				if item < 4 {
					shapes[block][0] = append(shapes[block][0], unpack_shape_code(line))
					areas[block] += strings.Count(line, "#")
				}
			}
		} else {
			parts := strings.Split(line, " ")
			nums := strings.Split(parts[0], "x")
			width, _ := strconv.Atoi(nums[0])
			length, _ := strconv.Atoi(strings.TrimRight(nums[1], ":"))
			presents := []int{}
			for i := 1; i < len(parts); i++ {
				n, _ := strconv.Atoi(parts[i])
				presents = append(presents, n)
			}
			regions = append(regions, region{width, length, presents})
		}
		line_num++
	}

	for i, sh := range shapes {
		shapes[i] = set_shape_variants(sh)
	}

	// Part 1
	// Strategy
	// Try to fit shape as low vertically as possible.
	// Then try to move the shape to the left.
	// Repeat for various combinations of shapes.

	// But this simple check of area works:

	total := 0
	for _, r := range regions {
		region_area := r.width * r.length
		min_present_area := 0
		for i, n := range r.presents {
			min_present_area += areas[i] * n
		}
		if region_area >= min_present_area {
			total++
		}
	}
	fmt.Printf("Part 1 total = %d\n", total)
}

func presentsFit(r region, shapes []shape) bool {
	// Create grid
	grid := make([][]bool, r.length)
	for i := range r.length {
		grid[i] = make([]bool, r.width)
	}

	// Generate a list of shape indices that must be placed in the grid
	slist := []int{}
	for i, n := range r.presents {
		if n > 0 {
			for range n {
				slist = append(slist, i)
			}
		}
	}

	// Get permutations
	perms := [][]int{}
	heapPermutation(slist, len(slist), &perms)

	for _, p := range perms {
		for _, i := range p {
			fmt.Println(i)
		}
	}

	//var func addshape(pidx, )
	return true
}

func heapPermutation(a []int, size int, perms *[][]int) {
	if size == 1 {
		if !slices.ContainsFunc(*perms, func(s []int) bool {
			return slices.Equal(s, a)
		}) {
			b := make([]int, len(a))
			copy(b, a)
			*perms = append(*perms, b)
		}
	}
	for i := range size {
		heapPermutation(a, size-1, perms)
		last := a[size-1]
		if size%2 == 0 {
			a[size-1] = a[i]
			a[i] = last
		} else {
			a[size-1] = a[0]
			a[0] = last
		}
	}
}

func unpack_shape_code(line string) [3]int {
	base_shape := [3]int{}
	for i, ch := range line {
		if ch == '#' {
			base_shape[i] = 1
		}
	}
	return base_shape
}

func set_shape_variants(sh shape) shape {
	result := shape{}
	sh = append(sh, flip_h(sh[0]))
	sh = append(sh, flip_v(sh[0]))
	sh = append(sh, flip_h(sh[2]))
	sh = append(sh, get_rotated_shape(sh[0]))
	sh = append(sh, flip_h(sh[4]))
	sh = append(sh, flip_v(sh[4]))
	sh = append(sh, flip_h(sh[6]))

	// set duplicates to nil
	for i := range len(sh) - 1 {
		if sh[i] != nil {
			for j := i + 1; j < len(sh); j++ {
				if sh[j] != nil && sh[i][0] == sh[j][0] && sh[i][1] == sh[j][1] && sh[i][2] == sh[j][2] {
					sh[j] = nil
				}
			}
		}
	}
	for _, s := range sh {
		if s != nil {
			result = append(result, s)
		}
	}
	return result
}

func get_rotated_shape(a [][3]int) [][3]int {
	b := make([][3]int, 3)
	b[1][1] = a[1][1]
	for i := range 2 {
		b[2-i][0] = a[0][i]
		b[0][i] = a[i][2]
		b[i][2] = a[2][2-i]
		b[2][2-i] = a[2-i][0]
	}
	return b
}

func flip_h(a [][3]int) [][3]int {
	b := make([][3]int, 3)
	copy(b, a)
	b[0][0] = a[0][2]
	b[1][0] = a[1][2]
	b[2][0] = a[2][2]
	b[0][2] = a[0][0]
	b[1][2] = a[1][0]
	b[2][2] = a[2][0]
	return b
}

func flip_v(a [][3]int) [][3]int {
	b := make([][3]int, 3)
	copy(b, a)
	b[0][0] = a[2][0]
	b[0][1] = a[2][1]
	b[0][2] = a[2][2]
	b[2][0] = a[0][0]
	b[2][1] = a[0][1]
	b[2][2] = a[0][2]
	return b
}
