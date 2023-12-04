package day3

import (
	"adventofcode2023/pkg/utils"
	"fmt"
	"regexp"
	"strconv"
)

type searchRange struct {
	StartingPoint int
	EndPoint      int
	Number        string
}

func Day31(filepath string) (int, error) {
	srFormerLine := []searchRange{}

	numberRegex := regexp.MustCompile("[0-9]+")
	matcher := regexp.MustCompile("[^0-9.]+")
	formerLine := ""
	result := 0
	firstline := true
	utils.ReadFile(filepath, func(line string) {
		longLine := fmt.Sprintf("%s%s%s", ".", line, ".")

		found := numberRegex.FindAllString(longLine, -1)
		foundIndexes := numberRegex.FindAllStringIndex(longLine, -1)
		srCurrentLine := *findAllNumberInfo(found, foundIndexes, longLine)
		foundResult, leftOver := updateResult(longLine, &srCurrentLine, matcher)
		foundResult2 := 0
		foundResult3 := 0
		srCurrentLine = *leftOver
		if firstline {
			firstline = false
		} else {

			foundResult22, notUsable := updateResult(longLine, &srFormerLine, matcher)
			foundResult2 = foundResult22
			foundResult33, leftOver2 := updateResult(formerLine, &srCurrentLine, matcher)
			foundResult3 = foundResult33
			srCurrentLine = *leftOver2

			for _, notcounted := range *notUsable {
				println(notcounted.Number)
			}
		}
		//srFormerLine = *leftOver2

		result = result + foundResult + foundResult2 + foundResult3

		formerLine = longLine

		srFormerLine = srCurrentLine

	})
	return result, nil
}

func updateResult(line string, srCurrentLine *[]searchRange, matcher *regexp.Regexp) (int, *[]searchRange) {
	toBeRemoved := []searchRange{}
	unfoundSrCurrentLine := srCurrentLine
	result := 0
	for _, searchRange := range *srCurrentLine {
		substring := line[searchRange.StartingPoint:searchRange.EndPoint]
		if matcher.MatchString(substring) {
			n, _ := strconv.Atoi(searchRange.Number)
			result = result + n
			toBeRemoved = append(toBeRemoved, searchRange)
		}
	}

	for _, remove := range toBeRemoved {
		for i, value := range *unfoundSrCurrentLine {
			if value == remove {
				(*unfoundSrCurrentLine)[i] = (*unfoundSrCurrentLine)[len(*unfoundSrCurrentLine)-1]
				bla := (*unfoundSrCurrentLine)[:len(*unfoundSrCurrentLine)-1]
				unfoundSrCurrentLine = &bla
				break
			}
		}

	}
	return result, unfoundSrCurrentLine
}

func findAllNumberInfo(found []string, indexes [][]int, longLine string) *[]searchRange {
	ranges := []searchRange{}
	for i, foundNumber := range found {
		ranges = append(ranges, searchRange{indexes[i][0] - 1, indexes[i][1] + 1, foundNumber})
	}
	return &ranges
}
