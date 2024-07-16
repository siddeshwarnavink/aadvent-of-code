package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ExtractNumbers(str string) []int {
	var numbers []int

	str = strings.TrimSpace(str)
	nums := strings.Split(str, " ")

	for i := 0; i < len(nums); i++ {
		numStr := nums[i]

		if numStr != "" {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				panic(err)
			}

			numbers = append(numbers, num)
		}

	}

	return numbers
}

func main() {
	data, err := os.ReadFile("day4.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")
	points := 0

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		winning := make(map[int]bool)

		if line != "" {
			numbersStr := strings.Split(line, ":")[1]
			cardPoints := 0

			winningPart := strings.Split(numbersStr, "|")[0]
			numbersPart := strings.Split(numbersStr, "|")[1]

			winningNumbers := ExtractNumbers(winningPart)
			numbersInHand := ExtractNumbers(numbersPart)

			for j := 0; j < len(winningNumbers); j++ {
				winningNumber := winningNumbers[j]
				winning[winningNumber] = true
			}

			for j := 0; j < len(numbersInHand); j++ {
				number := numbersInHand[j]
				if winning[number] {
					cardPoints++
				}
			}

			if cardPoints > 2 {
				cardPoints *= 2
			}

			points += cardPoints
		}
	}

	fmt.Println("points=", points)
}
