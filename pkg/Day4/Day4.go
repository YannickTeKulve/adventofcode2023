package Day4

import (
	"adventofcode2023/pkg/utils"
	"math"
	"regexp"
	"strings"
)

type copy struct {
	value int
}

func Task2(filepath string) int {

	result := 0
	numberegex, _ := regexp.Compile("[0-9]+")
	cardvalues := [][]int{}
	utils.ReadFile(filepath, func(line string) {
		game := strings.Split(line, ":")[1]
		gameInfo := strings.Split(game, "|")
		winningValues := numberegex.FindAllString(gameInfo[0], -1)
		scratchValues := numberegex.FindAllString(gameInfo[1], -1)
		cardValue := 0
		for _, scratch := range scratchValues {

			for _, winningValue := range winningValues {
				if scratch == winningValue {
					cardValue++
				}
			}
		}
		cards := append([]int{}, cardValue)
		cardvalues = append(cardvalues, cards)

	})

	for i, value := range cardvalues {
		for _, cardValue := range value {
			firstIndex := i + 1
			lastIndex := firstIndex + cardValue
			length := len(cardvalues)
			if firstIndex >= length {
				firstIndex = length - 1
			}
			if lastIndex > length {
				lastIndex = length
			}
			for j, cards := range cardvalues[firstIndex:lastIndex] {
				cardvalues[firstIndex+j] = append(cards, cards[0])
			}
			result++
		}
	}

	return result
}

func Task1(filepath string) int {

	result := 0
	numberegex, _ := regexp.Compile("[0-9]+")
	utils.ReadFile(filepath, func(line string) {
		game := strings.Split(line, ":")[1]
		gameInfo := strings.Split(game, "|")
		winningValues := numberegex.FindAllString(gameInfo[0], -1)
		scratchValues := numberegex.FindAllString(gameInfo[1], -1)
		cardValue := 0
		for _, scratch := range scratchValues {

			for _, winningValue := range winningValues {
				if scratch == winningValue {
					cardValue++
				}
			}

		}

		if cardValue > 2 {
			cardValue = powInt(2, cardValue-1)

		}
		result = result + cardValue

	})
	return result
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
