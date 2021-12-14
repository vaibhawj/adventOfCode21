package day3

import (
	"fmt"
	"strconv"
	"strings"
)

// part1
func PowerConsumption(input []string) int64 {
	gammaStr := ""
	epsilonStr := ""
	for j := 0; j < len(input[0]); j++ {
		countOf1 := 0
		countOf0 := 0
		for i := 0; i < len(input); i++ {
			if input[i][j] == '0' {
				countOf0++
			} else {
				countOf1++
			}
		}
		if countOf1 > countOf0 {
			gammaStr += "1"
			epsilonStr += "0"
		} else {
			gammaStr += "0"
			epsilonStr += "1"
		}
	}
	gamma, _ := strconv.ParseInt(gammaStr, 2, 64)
	fmt.Println(gammaStr)
	epsilon, _ := strconv.ParseInt(epsilonStr, 2, 64)
	fmt.Println(epsilonStr)
	fmt.Println(gamma * epsilon)
	return gamma * epsilon
}

// part2
func LifeSupportRating(input []string) int64 {
	oxyGenRatingStr := ""
	CarbScrubRatingStr := ""
	for j := 0; j < len(input[0]); j++ {
		countOf1 := 0
		countOf0 := 0
		filteredInputs := filterMatching(input, oxyGenRatingStr)
		if len(filteredInputs) == 1 {
			oxyGenRatingStr = filteredInputs[0]
			break
		}
		for i := 0; i < len(filteredInputs); i++ {
			if filteredInputs[i][j] == '0' {
				countOf0++
			} else {
				countOf1++
			}
		}
		if countOf1 >= countOf0 {
			oxyGenRatingStr += "1"
		} else {
			oxyGenRatingStr += "0"
		}
	}

	for j := 0; j < len(input[0]); j++ {
		countOf1 := 0
		countOf0 := 0
		filteredInputs := filterMatching(input, CarbScrubRatingStr)
		if len(filteredInputs) == 1 {
			CarbScrubRatingStr = filteredInputs[0]
			break
		}
		for i := 0; i < len(filteredInputs); i++ {
			if filteredInputs[i][j] == '0' {
				countOf0++
			} else {
				countOf1++
			}
		}
		if countOf0 <= countOf1 {
			CarbScrubRatingStr += "0"
		} else {
			CarbScrubRatingStr += "1"
		}
	}
	oxy, _ := strconv.ParseInt(oxyGenRatingStr, 2, 64)
	fmt.Println(oxyGenRatingStr)
	carb, _ := strconv.ParseInt(CarbScrubRatingStr, 2, 64)
	fmt.Println(CarbScrubRatingStr)
	fmt.Println(oxy * carb)
	return oxy * carb
}

func filterMatching(input []string, prefix string) (ret []string) {
	for _, s := range input {
		if strings.HasPrefix(s, prefix) {
			ret = append(ret, s)
		}
	}
	return
}
