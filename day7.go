package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type splitter struct {
	id, pos int
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	start_pos := strings.IndexByte(lines[0], 'S')
	next_node_id := 0
	nodes := map[int][]int{}

	// Part 1
	beams := map[int][]int{}
	beams[start_pos] = append(beams[start_pos], next_node_id)
	next_node_id++
	num_rows := len(lines)
	num_splits := 0
	row := 2
	for row < num_rows {
		splitters, next := find_splitters(lines[row], nodes, next_node_id)
		next_node_id = next
		new_beams := map[int][]int{}
		for beam_pos, ancestor_node_ids := range beams {
			splitter_idx := slices.IndexFunc(splitters, func(s splitter) bool {
				return s.pos == beam_pos
			})
			if splitter_idx > -1 {
				spid := splitters[splitter_idx].id
				num_splits++
				// Assign ancestor node id to the beams
				new_beams[beam_pos-1] = append(new_beams[beam_pos-1], spid)
				new_beams[beam_pos+1] = append(new_beams[beam_pos+1], spid)
				// Add this node to child list of ancestor nodes
				for _, ancestor_node_id := range ancestor_node_ids {
					nodes[ancestor_node_id] = append(nodes[ancestor_node_id], spid)
				}
			} else {
				new_beams[beam_pos] = append(new_beams[beam_pos], ancestor_node_ids...)
			}
		}
		beams = new_beams
		row += 2
		if row == num_rows {
			// Add end node to the beams
			nodes[next_node_id] = []int{}
			for _, ancestor_node_ids := range beams {
				for _, ancestor_node_id := range ancestor_node_ids {
					nodes[ancestor_node_id] = append(nodes[ancestor_node_id], next_node_id)
				}
			}
		}
	}

	fmt.Printf("Part 1. Number of splits = %d\n", num_splits)

	// Part 2
	timeline_count := 0
	// Perform a DF traversal of the nodes starting at node 0
	stack := []int{0}

	for len(stack) > 0 {
		idx := len(stack) - 1
		node_id := stack[idx]
		stack = stack[:idx]
		children := nodes[node_id]
		if len(children) == 0 {
			timeline_count++
		}
		stack = append(stack, children...)
	}

	fmt.Printf("Part 2. Number of timelines = %d\n", timeline_count)
}

func find_splitters(str string, splitters map[int][]int, next_id int) ([]splitter, int) {
	ids := []splitter{}
	for idx, chr := range str {
		if chr == '^' {
			splitters[next_id] = []int{}
			ids = append(ids, splitter{id: next_id, pos: idx})
			next_id++
		}
	}
	return ids, next_id
}
