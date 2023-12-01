package day1

import (
	"adventofcode2023/pkg/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Day11(filepath string) (int, error) {
	result := 0
	utils.ReadFile(filepath, func(line string) {
		re := regexp.MustCompile("1|2|3|4|5|6|7|8|9")
		found := re.FindAllString(line, -1)
		first := found[0]
		last := found[len(found)-1]
		printed := fmt.Sprintf("%s%s", first, last)
		number, _ := strconv.Atoi(printed)
		result = result + number

	})
	return result, nil
}

func Day12(filepath string) (int, error) {
	mapping := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	result := 0
	utils.ReadFile(filepath, func(line string) {
		re := regexp.MustCompile("1|2|3|4|5|6|7|8|9")

		newLine := line
		for written, number := range mapping {
			replacement := fmt.Sprintf("%s%s%s", string(written[0]), number, string(written[len(written)-1]))
			newLine = strings.ReplaceAll(newLine, written, replacement)
		}

		found := re.FindAllString(newLine, -1)

		first := found[0]

		last := found[len(found)-1]
		printed := fmt.Sprintf("%s%s", first, last)
		number, _ := strconv.Atoi(printed)
		result = result + number

	})
	return result, nil
}
