package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isNumber(str string) bool {
	_, err := strconv.Atoi(str)

	if err != nil {
		return false
	}

	return true
}

func Reverse(str string) string {
	var result strings.Builder
	for i := len(str) - 1; i >= 0; i-- {
		result.WriteByte(str[i])
	}

	return result.String()
}

func main() {
	data, err := os.ReadFile("day3.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")
	sum := 0

	for i := 0; i < len(lines); i++ {
		line := lines[i]

		for j := 0; j < len(line); j++ {
			ch := string(line[j])

			// Check if a symbol
			if !isNumber(ch) && ch != "." {
				// left
				if j >= 1 && isNumber(string(line[j-1])) {
					k := j - 1
					var numStr strings.Builder

					for k >= 0 && isNumber(string(line[k])) {
						numStr.WriteByte(line[k])
						k--
					}

					num, err := strconv.Atoi(Reverse(numStr.String()))
					if err != nil {
						panic(err)
					}

					sum += num

					fmt.Printf("%s found in left of %s\n", string(line[j-1]), ch)
				}

				// right
				if j+2 <= len(line) && isNumber(string(line[j+1])) {
					k := j + 1
					var numStr strings.Builder

					for k <= len(line)-1 && isNumber(string(line[k])) {
						numStr.WriteByte(line[k])
						k++
					}

					num, err := strconv.Atoi(numStr.String())
					if err != nil {
						panic(err)
					}

					sum += num

					fmt.Printf("%s found in right of %s\n", string(line[j+1]), ch)
				}

				// top
				if i-1 >= 0 && isNumber(string(lines[i-1][j])) {
					// check if number tending to left
					if j-1 >= 0 && isNumber(string(lines[i-1][j-1])) {
						k := j
						var numStr strings.Builder

						for k >= 0 && isNumber(string(lines[i-1][k])) {
							numStr.WriteByte(lines[i-1][k])
							k--
						}

						num, err := strconv.Atoi(Reverse(numStr.String()))
						if err != nil {
							panic(err)
						}

						sum += num
					}

					// check if number tending to right
					if j+2 <= len(lines[i-1]) && isNumber(string(lines[i-1][j+1])) {
						k := j
						var numStr strings.Builder

						for k <= len(lines[i-1])-1 && isNumber(string(lines[i-1][k])) {
							numStr.WriteByte(lines[i-1][k])
							k++
						}

						num, err := strconv.Atoi(numStr.String())
						if err != nil {
							panic(err)
						}

						sum += num
					}

					fmt.Printf("%s found in top of %s\n", string(lines[i-1][j]), ch)
				}
			}
		}
	}

	fmt.Println("answer=", sum)
}
