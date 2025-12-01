package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input") // Load the input data from a file called "input"
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	password := 0
	num := 50 // Range of 0 .. 99
	for _, str := range lines {
		// Get number
		clicks, err := strconv.Atoi(str[1:])
		if err != nil {
			log.Fatalf("Error converting string to integer: %v", err)
		}

		if str[0] == 'L' {
			for range clicks {
				num -= 1
				if num == 0 {
					password += 1
				}
				if num == -1 {
					num = 99
				}
			}
		} else {
			for range clicks {
				num += 1
				if num == 100 {
					password += 1
					num = 0
				}
			}
		}
		//fmt.Println(num)
	}
	fmt.Printf("Password is: %d\n", password)
}

/*
This code is easy to understand and the functionality
is easy to debug. It would lend itself well to driving an animated display of the solution.

It is not efficient though. I tried to write an efficient solution but did not get the correct
answer with it. An efficient solution would apply all of the increments in each step in one go,
but this is difficult for backwards steps involving negative values and multiple rotations.

Have to consider starting from 0 and multiple rotations in either direction.
/*
