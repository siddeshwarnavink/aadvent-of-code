package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("day1.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")
	sum := 0

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		var first, last *int

		for j := 0; j < len(line); j++ {
      val := int(line[j] - '0')

			if val <= 9 {
				if first == nil {
					first = &val
				}
				last = &val
			}
		}

		if first != nil && last != nil {
			num := (10 * *first) + *last
			sum += num
		}
	}

	fmt.Println(sum)
}
