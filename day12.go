package main // Incomplete

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"
)

type region struct {
	width, length int
	presents      []int
}

type shape [][][]int

func main() {
	file, err := os.Open("test")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	shapes := []shape{}
	regions := []region{}
	scanner := bufio.NewScanner(file)
	line_num := 0
	for scanner.Scan() {
		line := scanner.Text()
		block := line_num / 5
		item := line_num % 5
		if block < 6 {
			if item == 0 {
				shapes = append(shapes, shape{})
			} else {
				if item < 4 {
					shapes[block][0] = append(shapes[block][0], unpack_shape_code(line))
					set_shape_variants(shapes[block])
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
}

func unpack_shape_code(line string) []int {
	base_shape := make([]int, 3)
	for i, ch := range line {
		if ch == '#' {
			base_shape[i] = 1
		}
	}
	return base_shape
}

func set_shape_variants(sh shape) {
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
				if slices.Equal(sh[i][0], sh[j][0]) && slices.Equal(sh[i][1], sh[j][1]) && slices.Equal(sh[i][2], sh[j][2]) {
					sh[j] = nil
				}
			}
		}
	}
}

func get_rotated_shape(a [][]int) [][]int {
	b := make([][]int, 3)
	for range 3 {
		b = append(b, make([]int, 3))
	}
	for i := range 3 {
		b[2-i][0] = a[0][i]
		b[0][i] = a[i][2]
		b[i][2] = a[2][2-i]
		b[2][i] = a[i][0]
	}
	return b
}

func flip_h(a [][]int) [][]int {
	b := make([][]int, 3)
	for range 3 {
		b = append(b, make([]int, 3))
	}
	for range 3 {
		b[0][0] = a[0][2]
		b[1][0] = a[1][2]
		b[2][0] = a[2][2]
		b[0][2] = a[0][0]
		b[1][2] = a[1][0]
		b[2][2] = a[2][0]
	}
	return b
}

func flip_v(a [][]int) [][]int {
	b := make([][]int, 3)
	for range 3 {
		b = append(b, make([]int, 3))
	}
	for range 3 {
		b[0][0] = a[2][0]
		b[0][1] = a[2][1]
		b[0][2] = a[2][2]
		b[2][0] = a[0][0]
		b[2][1] = a[0][1]
		b[2][2] = a[0][2]
	}
	return b
}
