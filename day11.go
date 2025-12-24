package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	strs := map[string][]string{}
	keys := map[string]int{}
	ki := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, ": ")
		strs[items[0]] = strings.Split(items[1], " ")
		keys[items[0]] = ki
		ki++
	}
	strs["out"] = []string{}
	keys["out"] = ki

	devices := make([][]int, len(strs))
	for k, v := range strs {
		dests := make([]int, len(v))
		for i, d := range v {
			dests[i] = keys[d]
		}
		devices[keys[k]] = dests
	}
	n := countPaths(keys["you"], devices, keys["out"])
	fmt.Printf("Part 1 number of paths = %d\n", n)

	n = countPaths(keys["svr"], devices, keys["fft"])
	n *= countPaths(keys["fft"], devices, keys["dac"])
	n *= countPaths(keys["dac"], devices, keys["out"])
	fmt.Printf("Part 2 number of paths = %d\n", n)

}

func countPaths(start int, devices [][]int, goal int) int {
	memo := make(map[int]int)
	var dfs func(int) int
	dfs = func(pos int) int {
		if pos == goal {
			return 1
		}
		if val, ok := memo[pos]; ok {
			return val
		}
		total := 0
		for _, next := range devices[pos] {
			total += dfs(next)
		}
		memo[pos] = total
		return total
	}
	return dfs(start)
}
