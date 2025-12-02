package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("day2.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")
	sum := 0

	for i := 0; i < len(lines); i++ {
		line := lines[i]

		red := 0
		green := 0
		blue := 0

		if line != "" {
			gameIdStr := strings.Split(line, " ")[1][0:1]
			gameId, err := strconv.Atoi(gameIdStr)
			if err != nil {
				panic(err)
			}

			rounds := strings.Split(line[8:len(line)], ";")

			for j := 0; j < len(rounds); j++ {
				balls := strings.Split(rounds[j], ",")

				for k := 0; k < len(balls); k++ {
					ball := strings.TrimSpace(balls[k])

					count, err := strconv.Atoi(strings.Split(ball, " ")[0])
					if err != nil {
						panic(err)
					}

					ballName := strings.Split(ball, " ")[1]

					switch ballName {
					case "red":
						red += count
						break

					case "green":
						green += count
						break

					case "blue":
						blue += count
						break

					default:
						panic("are you alright baby?")
					}
				}
			}

			if red < 12 && green < 13 && blue < 14 {
				sum += gameId
			}
		}
	}

  fmt.Println(sum)
}
