package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	contentBytes, err := os.ReadFile("input")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	ranges := string(contentBytes)
	p1_sum := 0
	p2_sum := 0
	for _, id_range := range strings.Split(ranges, ",") {
		num_pair := strings.Split(id_range, "-")
		start_id, _ := strconv.Atoi(num_pair[0])
		end_id, _ := strconv.Atoi(num_pair[1])
		for i := start_id; i <= end_id; i++ {
			num_digits := int(math.Log10(float64(i))) + 1

			// Part 1
			// Check for an even number of digits
			if num_digits%2 == 0 {
				div := int(math.Pow10(num_digits / 2))
				if i/div == i%div {
					p1_sum += i
				}
			}
			// End of Part 1

			// Part 2
			ndigits := num_digits / 2
			for ndigits > 0 {
				number := i
				if num_digits%ndigits == 0 {
					steps := num_digits / ndigits
					// Shift digits right and compare chunks
					var last_chunk int
					div := int(math.Pow10(ndigits))
					invalid := true
					for j := 0; j < steps; j++ {
						new_chunk := number % div
						number /= div
						if j > 0 && last_chunk != new_chunk {
							invalid = false
							break
						}
						last_chunk = new_chunk
					}
					if invalid {
						p2_sum += i
						break
					}
				}
				ndigits -= 1
			}
		}
	}
	fmt.Println(p1_sum, p2_sum)
}
