package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type button []int

type machine struct {
	lights   []bool
	buttons  []button
	joltages []int
}

func main() {
	file, err := os.Open("test")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var machines []machine
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		m := machine{}
		for _, chr := range parts[0] {
			switch chr {
			case '.':
				m.lights = append(m.lights, false)
			case '#':
				m.lights = append(m.lights, true)
			}
		}
		// Observe that switch digits appear to be single digits
		for i := 1; i < len(parts)-1; i++ {
			b := button{}
			for _, chr := range parts[i] {
				num, err := strconv.Atoi(string(chr))
				if err == nil {
					b = append(b, num)
				}
			}
			m.buttons = append(m.buttons, b)
		}
		jolts := strings.SplitSeq(strings.Trim(parts[len(parts)-1], "{}"), ",")
		for jolt := range jolts {
			jv, _ := strconv.Atoi(jolt)
			m.joltages = append(m.joltages, jv)
		}
		machines = append(machines, m)
	}
	fmt.Println(machines)
}
