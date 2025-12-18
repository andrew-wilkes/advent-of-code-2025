package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type device struct {
	id       int
	dac, fft bool
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	strs := map[string][]string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, ": ")
		strs[items[0]] = strings.Split(items[1], " ")
	}
	strs["out"] = []string{}

	// Convert strings to ints
	devices := make([][]device, len(strs))
	keys := map[string]int{}
	ki := 0
	for key, _ := range strs {
		keys[key] = ki
		ki++
	}
	for k, v := range strs {
		dests := make([]device, len(v))
		for i, d := range v {
			dests[i] = device{id: keys[d]}
		}
		devices[keys[k]] = dests
	}
	//n := propagate_dfs(keys["you"], devices, []int{}, 0, keys["out"])
	n := propagate_bfs(device{id: keys["you"]}, devices, keys["out"])
	fmt.Printf("Part 1 number of paths = %d\n", n)

	//n = propagate2("svr", strs, []string{}, 0, false, false)
	n = propagate_bfs2(device{id: keys["svr"]}, devices, keys["out"], keys["dac"], keys["fft"])
	fmt.Printf("Part 2 number of paths = %d\n", n)

}

func propagate_dfs(from device, devices [][]device, path []device, n, target int) int {
	path = append(path, from)
	for _, dest := range devices[from.id] {
		if dest.id == target {
			return n + 1
		} else {
			if !slices.Contains(path, dest) {
				n = propagate_dfs(dest, devices, path, n, target)
			}
		}
	}
	return n
}

// This finds the shortest path, not all paths, if visited nodes are tracked
func propagate_bfs(from device, devices [][]device, target int) int {
	n := 0
	visited := []device{}
	queue := []device{from}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.id == target {
			n++
		} else {
			if !slices.Contains(visited, node) {
				//visited = append(visited, node)
				queue = append(queue, devices[node.id]...)
			}
		}
	}
	return n
}

func propagate_bfs2(from device, devices [][]device, target, dac, fft int) int {
	n := 0
	queue := []device{from}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		switch node.id {
		case target:
			if node.dac && node.fft {
				n++
			}
		case dac:
			node.dac = true
		case fft:
			node.fft = true
		}
		idx := len(queue)
		queue = append(queue, devices[node.id]...)
		if node.dac || node.fft {
			for i := idx; i < len(queue); i++ {
				queue[i].dac = node.dac
				queue[i].fft = node.fft
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
