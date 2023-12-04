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

type validgear struct {
	linenr int
	index  []int
}

func Day31task2(filepath string) (int, error) {
	numberRegex := regexp.MustCompile("[0-9]+")
	var lines []string
	var gearLines = map[int][][]int{}
	validGears := []validgear{}
	linenumber := 0
	result := 0
	utils.ReadFile(filepath, func(line string) {
		numberRegex := regexp.MustCompile("[\\*]+")
		gears := numberRegex.FindAllStringIndex(line, -1)
		if len(gears) > 0 {
			gearLines[linenumber] = gears
		}
		lines = append(lines, line)
		linenumber++
	})

	for i, gears := range gearLines {
		for _, gear := range gears {
			onLineAbove := 0
			onLineUnder := 0
			if (i - 1) >= 0 {
				onLineAbove = len(numberRegex.FindAllString(lines[i-1][gear[0]-1:gear[1]+1], -1))
			}
			onLineItself := len(numberRegex.FindAllString(lines[i][gear[0]-1:gear[1]+1], -1))
			if (i + 1) < len(lines) {
				onLineUnder = len(numberRegex.FindAllString(lines[i+1][gear[0]-1:gear[1]+1], -1))
			}
			found := onLineAbove + onLineUnder + onLineItself
			if found == 2 {
				validGears = append(validGears, validgear{i, gear})

			}
			println(found)
		}
	}

	for _, valid := range validGears {
		grid := map[int][][]int{}

		if (valid.linenr - 1) >= 0 {
			grid[valid.linenr-1] = numberRegex.FindAllStringIndex(lines[valid.linenr-1], -1)
		}
		grid[valid.linenr] = numberRegex.FindAllStringIndex(lines[valid.linenr], -1)
		if (valid.linenr + 1) < len(lines) {
			grid[valid.linenr+1] = numberRegex.FindAllStringIndex(lines[valid.linenr+1], -1)
		}

		gearResult := 0
		for linenr, line := range grid {
			for _, numberIndex := range line {
				println(fmt.Sprintf("checking: %s", lines[linenr][numberIndex[0]:numberIndex[1]]))
				if (numberIndex[0] <= valid.index[0] && numberIndex[1] >= valid.index[0]) ||
					(numberIndex[0] <= (valid.index[0]-1) && numberIndex[1] > (valid.index[0]-1)) ||
					(numberIndex[0] <= (valid.index[0]+1) && numberIndex[1] > (valid.index[0]+1)) {
					println(fmt.Sprintf("%d - %d", linenr, valid.index[0]))
					println(lines[linenr][numberIndex[0]:numberIndex[1]])
					value, _ := strconv.Atoi(lines[linenr][numberIndex[0]:numberIndex[1]])
					if gearResult == 0 {
						gearResult = value
					} else {
						gearResult = gearResult * value
					}
				}
			}
		}
		println(fmt.Sprintf("gearresult: %d", gearResult))
		result = gearResult + result

	}

	return result, nil
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
