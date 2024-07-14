package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Number struct {
	num    int
	numLen int
	x      int
	y      int
	next   *Number
	prev   *Number
}

func (n *Number) Append(number *Number) {
	for n.next != nil {
		n = n.next
	}
	number.prev = n
	n.next = number
}

func (n *Number) Remove() {
	if n.prev != nil {
		n.prev.next = n.next
	}

	if n.next != nil {
		n.next.prev = n.prev
	}

	n.next = nil
	n.prev = nil
}

func isNumStr(str string) bool {
	_, err := strconv.Atoi(str)

	if err != nil {
		return false
	}

	return true
}

// TODO Complete this
func checkAdjecent(number *Number, x int, y int) bool {
	// // y axis
	// if number.y-number.numLen-1 == y+1 {
	// 	return true
	// }
	// if number.y == y-1 {
	// 	return true
	// }

	// for i := 0; i < number.numLen; i++ {
	// 	// x axis
	// 	if number.x-i == x+1 {
	// 		return true
	// 	}
	// 	if number.x-i == x-1 {
	// 		return true
	// 	}
	// }

  fmt.Printf("checking for %d\n", number.num)

	if number.x == x+1 || number.x == x-1 {
		return true
	}

	return false
}

func main() {
	data, err := os.ReadFile("day3.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")
	isPrevInt := false
	var head *Number
	var num strings.Builder

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		isPrevInt = false

		for j := 0; j < len(line); j++ {
			ch := string(line[j])

			if isNumStr(ch) {
				isPrevInt = true
				num.WriteString(ch)
			} else if isPrevInt {
				chInt, err := strconv.Atoi(num.String())
				if err != nil {
					panic(err)
				}

				number := Number{
					num:    chInt,
					numLen: len(num.String()),
					x:      i,
					y:      j,
				}
				if head == nil {
					head = &number
				} else {
					head.Append(&number)
				}

				isPrevInt = false
				num = strings.Builder{}
			}

			if !isNumStr(ch) && ch != "." && head != nil {
				ptr := head

				for ptr.next != nil {
					// TODO: Complete this
					if checkAdjecent(ptr, i, j) {
						fmt.Printf("%d for char %s\n", ptr.num, ch)
						// push number to sum
					}
					ptr = ptr.next
				}
			}
		}
	}
}
