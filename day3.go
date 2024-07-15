package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type VisitedList map[int]map[int]bool

func (v VisitedList) Set(x, y int, value bool) {
	if v[x] == nil {
		v[x] = make(map[int]bool)
	}
	v[x][y] = value
}

func (v VisitedList) Get(x, y int) bool {
	if row, exists := v[x]; exists {
		_, exists := row[y]
		return exists
	}
	return false
}

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

func ExtractNumberTendingLeft(line *[]string, start int, height int, visited *VisitedList) int {
	k := start
	var numStr strings.Builder

	if visited.Get(height, k) {
		return 0
	}

	for k >= 0 && isNumber((*line)[k]) {
		visited.Set(height, k, true)
		numStr.WriteString((*line)[k])
		k--
	}

	num, err := strconv.Atoi(Reverse(numStr.String()))
	if err != nil {
		panic(err)
	}

	return num
}

func ExtractNumberTendingRight(line *[]string, start int, height int, visited *VisitedList) int {
	k := start
	var numStr strings.Builder

	if visited.Get(height, k) {
		return 0
	}

	for k <= len(*line)-1 && isNumber((*line)[k]) {
		visited.Set(height, k, true)
		numStr.WriteString((*line)[k])
		k++
	}

	num, err := strconv.Atoi(numStr.String())
	if err != nil {
		panic(err)
	}

	return num
}

func main() {
	data, err := os.ReadFile("day3.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")
	visited := make(VisitedList)
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
					sum += ExtractNumberTendingLeft(&lineRef, j-1, i, &visited)

					fmt.Printf("%s found in left of %s\n", string(line[j-1]), ch)
				}

				// right
				if j+2 <= len(line) && isNumber(string(line[j+1])) {
					lineRef := strings.Split(line, "")
					sum += ExtractNumberTendingRight(&lineRef, j+1, i, &visited)

					fmt.Printf("%s found in right of %s\n", string(line[j+1]), ch)
				}

				// top
				if i-1 >= 0 && isNumber(string(lines[i-1][j])) {
					// check if number tending to left
					if j-1 >= 0 && isNumber(string(lines[i-1][j-1])) {
						lineRef := strings.Split(lines[i-1], "")
						sum += ExtractNumberTendingLeft(&lineRef, j, i-1, &visited)
					} else if j+2 <= len(lines[i-1]) && isNumber(string(lines[i-1][j+1])) { // check if number tending to right
						lineRef := strings.Split(lines[i-1], "")
						sum += ExtractNumberTendingRight(&lineRef, j, i-1, &visited)
					}

					fmt.Printf("%s found in top of %s\n", string(lines[i-1][j]), ch)
				}

				// bottom
				if i+2 <= len(lines) && isNumber(string(lines[i+1][j])) {
					// check if number tending to left
					if j-1 >= 0 && isNumber(string(lines[i+1][j-1])) {
						lineRef := strings.Split(lines[i+1], "")
						sum += ExtractNumberTendingLeft(&lineRef, j, i+1, &visited)
					} else if j+2 <= len(lines[i+1]) && isNumber(string(lines[i+1][j+1])) { // check if number tending to right
						lineRef := strings.Split(lines[i+1], "")
						sum += ExtractNumberTendingRight(&lineRef, j, i+1, &visited)
					}

					fmt.Printf("%s found in bottom of %s\n", string(lines[i+1][j]), ch)
				}

				// diagonal top-left
				if i-1 >= 0 && isNumber(string(lines[i-1][j-1])) {
					// check if number tending to left
					if j-1 >= 0 && isNumber(string(lines[i-1][j-2])) {
						lineRef := strings.Split(lines[i-1], "")
						sum += ExtractNumberTendingLeft(&lineRef, j-1, i-1, &visited)
					} else if j+2 <= len(lines[i-1]) && isNumber(string(lines[i-1][j])) { // check if number tending to right
						lineRef := strings.Split(lines[i-1], "")
						sum += ExtractNumberTendingRight(&lineRef, j-1, i-1, &visited)
					}

					fmt.Printf("%s found in top-left of %s\n", string(lines[i-1][j-1]), ch)
				}

				// diagonal top-right
				if i-1 >= 0 && isNumber(string(lines[i-1][j+1])) {
					// check if number tending to left
					if j-1 >= 0 && isNumber(string(lines[i-1][j])) {
						lineRef := strings.Split(lines[i-1], "")
						sum += ExtractNumberTendingLeft(&lineRef, j+1, i-1, &visited)
					} else if j+2 <= len(lines[i-1]) && isNumber(string(lines[i-1][j+2])) { // check if number tending to right
						lineRef := strings.Split(lines[i-1], "")
						sum += ExtractNumberTendingRight(&lineRef, j+1, i-1, &visited)
					}

					fmt.Printf("%s found in top-right of %s\n", string(lines[i-1][j+1]), ch)
				}

				// diagonal bottom-left
				if i+2 <= len(lines) && isNumber(string(lines[i+1][j-1])) {
					// check if number tending to left
					if j-1 >= 0 && isNumber(string(lines[i+1][j-2])) {
						lineRef := strings.Split(lines[i+1], "")
						sum += ExtractNumberTendingLeft(&lineRef, j-1, i+1, &visited)
					} else if j+2 <= len(lines[i+1]) && isNumber(string(lines[i+1][j])) { // check if number tending to right
						lineRef := strings.Split(lines[i+1], "")
						sum += ExtractNumberTendingRight(&lineRef, j-1, i+1, &visited)
					}

					fmt.Printf("%s found in bottom-left of %s\n", string(lines[i+1][j-1]), ch)
				}

        // TODO: similary check bottom-right diagonal
			}
		}
	}

	fmt.Println("sum=", sum)
}
