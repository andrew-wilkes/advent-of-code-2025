package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

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

	password := 0
	num := 50 // Range of 0 .. 99
	for _, str := range lines {
		// Get number
		clicks, err := strconv.Atoi(str[1:])
		if err != nil {
			log.Fatalf("Error converting string to integer: %v", err)
		}

		if str[0] == 'L' {
			/*
				for range clicks {
					num -= 1
					if num == 0 {
						password += 1
					}
					if num == -1 {
						num = 99
					}
				}
			*/
			num_loops := 0
			if num > 0 && clicks >= num {
				num_loops++
			}
			num -= clicks
			num_loops -= num / 100
			num = (100 + num%100) % 100
			password += num_loops
		} else {
			/*
				for range clicks {
					num += 1
					if num == 100 {
						password += 1
						num = 0
					}
				}
			*/
			num += clicks
			num_loops := num / 100
			num %= 100
			password += num_loops
		}
		//fmt.Println(num)
	}
	fmt.Printf("Password is: %d\n", password)
}
