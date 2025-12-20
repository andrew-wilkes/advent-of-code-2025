package main // I tried out lots of ideas for part 2 but ultimately failed to solve it.

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
	file, err := os.Open("test2")
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

	devices := make([][]device, len(strs))
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
	n = bidirectional_bfs(device{id: keys["svr"]}, devices, keys["out"], keys["dac"], keys["fft"])
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

func bidirectional_bfs(from device, devices [][]device, target, dac, fft int) int {
	// Form the reverse path list
	d2 := make([][]device, len(devices))
	for n := range len(d2) {
		src := []device{}
		for i := range len(d2) {
			for _, d := range devices[i] {
				if d.id == n {
					src = append(src, device{id: i})
				}
			}
		}
		d2[n] = src
	}
	n := 0
	qa := []device{from}
	qb := []device{{id: target}}
	for len(qa) > 0 && len(qb) > 0 {
		if len(qb) > len(qa) {
			node := qa[0]
			qa = qa[1:]
			if node.id < 0 {
				continue // Ignore it
			}
			i := slices.Index(qb, node)
			if i > -1 {
				if (node.dac || qb[i].dac) && (node.fft || qb[i].fft) {
					n++
					qb = slices.Delete(qb, i, i+1)
					continue
				}
				if (node.dac == qb[i].dac) && (node.fft == qb[i].fft) {
					// Remove matching node from qb
					qb = slices.Delete(qb, i, i+1)
					continue // Skip further processing of this node
				}
				if !(qb[i].dac || qb[i].fft) {
					qa = slices.Delete(qb, i, i+1)
				}
				if !(node.dac || node.fft) {
					continue
				}
			}
			switch node.id {
			case dac:
				node.dac = true
			case fft:
				node.fft = true
			}
			idx := len(qa)
			qa = append(qa, devices[node.id]...)
			if node.dac || node.fft {
				for i := idx; i < len(qa); i++ {
					qa[i].dac = node.dac
					qa[i].fft = node.fft
				}
			}
		} else {
			node := qb[0]
			qb = qb[1:]
			if node.id < 0 {
				continue // Ignore it
			}
			i := slices.IndexFunc(qa, func(dev device) bool {
				return dev.id == node.id
			})
			if i > -1 {
				if (node.dac || qa[i].dac) && (node.fft || qa[i].fft) {
					n++
					qa = slices.Delete(qa, i, i+1)
					continue
				}
				if (node.dac == qa[i].dac) && (node.fft == qa[i].fft) {
					// Remove matching node from qa
					qa = slices.Delete(qa, i, i+1)
					continue // Skip further processing of this node
				}
				if !(qa[i].dac || qa[i].fft) {
					qa = slices.Delete(qa, i, i+1)
				}
				if !(node.dac || node.fft) {
					continue
				}
			}
			switch node.id {
			case dac:
				node.dac = true
			case fft:
				node.fft = true
			}
			idx := len(qb)
			qb = append(qb, d2[node.id]...)
			if node.dac || node.fft {
				for i := idx; i < len(qb); i++ {
					qb[i].dac = node.dac
					qb[i].fft = node.fft
				}
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
