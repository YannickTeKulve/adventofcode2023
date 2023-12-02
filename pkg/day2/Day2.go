package day2

import (
	"adventofcode2023/pkg/utils"
	"strconv"
	"strings"
)

func Day21(filepath string) (int, error) {
	mapping := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	result := 0
	utils.ReadFile(filepath, func(line string) {
		game, cubesCombosArray := getCubeCombos(line)
		for _, hand := range cubesCombosArray {
			handSplitted := strings.Split(hand, ",")
			if checkHandInvalid(handSplitted, mapping) {
				return
			}
		}

		gamenumber, _ := strconv.Atoi(strings.TrimPrefix(game, "Game "))
		result = result + gamenumber

	})
	return result, nil
}

func Day22(filepath string) (int, error) {
	result := 0
	utils.ReadFile(filepath, func(line string) {
		mapping := map[string]int{}
		_, cubesCombosArray := getCubeCombos(line)
		for _, hand := range cubesCombosArray {
			handSplitted := strings.Split(hand, ",")
			findPowerOfGame(handSplitted, &mapping)

		}

		power := 0
		for _, value := range mapping {
			if power == 0 {
				power = value
			} else {
				power = power * value
			}
		}
		result = result + power

	})
	return result, nil
}

func findPowerOfGame(cubesCombosArray []string, mapping *map[string]int) {
	cubesCombosMap := mapCubeCombos(cubesCombosArray)
	for color, valueHand := range cubesCombosMap {
		bagValue, ok := (*mapping)[color]
		if !ok {
			(*mapping)[color] = valueHand
			continue
		}
		if valueHand > bagValue {
			(*mapping)[color] = valueHand
		}
	}
}

func checkHandInvalid(cubesCombosArray []string, mapping map[string]int) bool {
	if len(cubesCombosArray) > 3 {
		return true
	}
	cubesCombosMap := mapCubeCombos(cubesCombosArray)
	for color, valueHand := range cubesCombosMap {
		bagValue, ok := mapping[color]
		if !ok {
			return true
		}
		if valueHand > bagValue {
			return true
		}
	}
	return false
}

func mapCubeCombos(cubecombos []string) map[string]int {
	comboMap := map[string]int{}
	for _, combo := range cubecombos {
		splitted := strings.Split(strings.TrimPrefix(combo, " "), " ")
		number, _ := strconv.Atoi(splitted[0])
		comboMap[splitted[1]] = number
	}
	return comboMap
}

func getCubeCombos(line string) (string, []string) {
	splitted := strings.Split(line, ":")
	cubeCombos := strings.Split(splitted[1], ";")
	return splitted[0], cubeCombos
}
