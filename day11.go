package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	devices := map[string][]string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, ": ")
		devices[items[0]] = strings.Split(items[1], " ")
	}
	n := propagate("you", devices, []string{}, 0)
	fmt.Printf("Part 1 number of paths = %d\n", n)

	n = propagate2("svr", devices, []string{}, 0, false, false)
	fmt.Printf("Part 2 number of paths = %d\n", n)

}

func propagate(from string, devices map[string][]string, path []string, n int) int {
	path = append(path, from)
	for _, dest := range devices[from] {
		if dest == "out" {
			return n + 1
		} else {
			if !slices.Contains(path, dest) {
				n = propagate(dest, devices, path, n)
			}
		}
	}
	return n
}

func propagate2(from string, devices map[string][]string, path []string, n int, dac, fft bool) int {
	path = append(path, from)
	for _, dest := range devices[from] {
		switch dest {
		case "out":
			if dac && fft {
				n++
			}
			return n
		case "dac":
			dac = true
		case "fft":
			fft = true
		}
		if !slices.Contains(path, dest) {
			n = propagate2(dest, devices, path, n, dac, fft)
		}
	}
	return n
}
