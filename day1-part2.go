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
