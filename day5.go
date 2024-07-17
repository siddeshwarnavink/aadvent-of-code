package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type SeedData struct {
	seed        int
	soil        int
	fertilizer  int
	water       int
	light       int
	temperature int
	humidity    int
	location    int
}

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
	var MAPPING_MODES = [...]string{"seed-to-soil", "soil-to-fertilizer", "fertilizer-to-water", "water-to-light", "light-to-temperature", "temperature-to-humidity", "humidity-to-location"}

	data, err := os.ReadFile("day5.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")
	var currentMode string
	var seedsToDisplay []int

	seedToSoilMap := make(map[int]int)
	soilToFertilizers := make(map[int]int)
	fertilizerToWater := make(map[int]int)
	waterToLight := make(map[int]int)
	lightToTemperature := make(map[int]int)
	temperatureToHumidity := make(map[int]int)
	humidityToLocation := make(map[int]int)

	for i := 0; i < len(lines); i++ {
		line := lines[i]

		if line == "" {
			currentMode = ""
			continue
		}

		// seeds to display
		if strings.Contains(line, "seeds:") {
			seedsStr := line[6:len(line)]
			seedsToDisplay = ExtractNumbers(seedsStr)
		}

		// listen change in mode
		if currentMode == "" {
			for j := 0; j < len(MAPPING_MODES); j++ {
				mode := MAPPING_MODES[j]
				if strings.Contains(line, mode) {
					currentMode = mode
				}
			}
		} else { // handle different mapping
			var myMap *(map[int]int)

			switch currentMode {
			case "seed-to-soil":
				myMap = &seedToSoilMap
				break
			case "soil-to-fertilizer":
				myMap = &soilToFertilizers
				break
			case "fertilizer-to-water":
				myMap = &fertilizerToWater
				break
			case "water-to-light":
				myMap = &waterToLight
				break
			case "light-to-temperature":
				myMap = &lightToTemperature
				break
			case "temperature-to-humidity":
				myMap = &temperatureToHumidity
				break
			case "humidity-to-location":
				myMap = &humidityToLocation
				break
			default:
				panic("something wrong baby?")
			}

			numbers := ExtractNumbers(line)
			for j := 0; j < numbers[2]; j++ {
				(*myMap)[numbers[1]+j] = numbers[0] + j
			}
		}
	}

	dataToDisplay := make([]SeedData, len(seedsToDisplay))

	// process data to display
	for i := 0; i < len(seedsToDisplay); i++ {
		var myData SeedData

		seed := seedsToDisplay[i]
		myData.seed = seed

		if soil, daijobu := seedToSoilMap[seed]; daijobu {
			myData.soil = soil
		} else {
			myData.soil = i
		}

		if fertilizer, daijobu := soilToFertilizers[myData.soil]; daijobu {
			myData.fertilizer = fertilizer
		} else {
			myData.fertilizer = i
		}

		if water, daijobu := fertilizerToWater[myData.fertilizer]; daijobu {
			myData.water = water
		} else {
			myData.water = i
		}

		if light, daijobu := waterToLight[myData.water]; daijobu {
			myData.light = light
		} else {
			myData.light = i
		}

		if temperature, daijobu := lightToTemperature[myData.light]; daijobu {
			myData.temperature = temperature
		} else {
			myData.temperature = i
		}

		if humidity, daijobu := temperatureToHumidity[myData.temperature]; daijobu {
			myData.humidity = humidity
		} else {
			myData.humidity = i
		}

		if location, daijobu := humidityToLocation[myData.humidity]; daijobu {
			myData.location = location
		} else {
			myData.location = i
		}

		dataToDisplay[i] = myData
	}

	// TODO: Quick sort dataToDisplay based on location

	// display the data
	for i := 0; i < len(dataToDisplay); i++ {
		data := dataToDisplay[i]
		fmt.Printf("Seed number %d corresponds to soil number %d.\n", data.seed, data.soil)
	}
}
