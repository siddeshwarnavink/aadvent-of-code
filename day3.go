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

func ExtractNumberTendingLeft(line *[]string, start int) int {
	k := start
	var numStr strings.Builder

	for k >= 0 && isNumber((*line)[k]) {
		numStr.WriteString((*line)[k])
		k--
	}

	num, err := strconv.Atoi(Reverse(numStr.String()))
	if err != nil {
		panic(err)
	}

  fmt.Println("got", num)

	return num
}

func ExtractNumberTendingRight(line *[]string, start int) int {
	k := start
	var numStr strings.Builder

	for k <= len(*line)-1 && isNumber((*line)[k]) {
		numStr.WriteString((*line)[k])
		k++
	}

	num, err := strconv.Atoi(numStr.String())
	if err != nil {
		panic(err)
	}

  fmt.Println("got", num)

	return num
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
					lineRef := strings.Split(line, "")
					sum += ExtractNumberTendingLeft(&lineRef, j-1)

					fmt.Printf("%s found in left of %s\n", string(line[j-1]), ch)
				}

				// right
				if j+2 <= len(line) && isNumber(string(line[j+1])) {
					lineRef := strings.Split(line, "")
					sum += ExtractNumberTendingLeft(&lineRef, j+1)

					fmt.Printf("%s found in right of %s\n", string(line[j+1]), ch)
				}

				// top
				if i-1 >= 0 && isNumber(string(lines[i-1][j])) {
					// check if number tending to left
					if j-1 >= 0 && isNumber(string(lines[i-1][j-1])) {
						lineRef := strings.Split(lines[i-1], "")
						sum += ExtractNumberTendingLeft(&lineRef, j)
					}

					// check if number tending to right
					if j+2 <= len(lines[i-1]) && isNumber(string(lines[i-1][j+1])) {
						lineRef := strings.Split(lines[i-1], "")
						sum += ExtractNumberTendingRight(&lineRef, j)
					}

					fmt.Printf("%s found in top of %s\n", string(lines[i-1][j]), ch)
				}

				// bottom
				if i+2 <= len(lines) && isNumber(string(lines[i+1][j])) {
					// check if number tending to left
					if j-1 >= 0 && isNumber(string(lines[i+1][j-1])) {
						lineRef := strings.Split(lines[i+1], "")
						sum += ExtractNumberTendingLeft(&lineRef, j)
					}

					// check if number tending to right
					if j+2 <= len(lines[i+1]) && isNumber(string(lines[i+1][j+1])) {
						lineRef := strings.Split(lines[i+1], "")
						sum += ExtractNumberTendingRight(&lineRef, j)
					}

					fmt.Printf("%s found in bottom of %s\n", string(lines[i+1][j]), ch)
				}
			}
		}
	}

	fmt.Println("answer=", sum)
}
