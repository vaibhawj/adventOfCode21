package day10

import (
	"sort"
	"strings"
)

func SyntaxErrorScore(input string) int {
	lines := strings.Split(input, "\n")
	errScore := 0
	for _, line := range lines {
		err := ScanError(line)
		if err != "" {
			switch err {
			case ")":
				errScore += 3
			case "]":
				errScore += 57
			case "}":
				errScore += 1197
			case ">":
				errScore += 25137

			}
		}
	}

	return errScore
}

func ScanError(line string) string {
	stack := makeStack()
	for _, char := range strings.Split(line, "") {
		if isOpening(char) {
			stack.push(char)
		} else {
			poppedChar := stack.pop()
			if openCloseDict[poppedChar] != char {
				return char
			}
		}
	}
	return ""
}

var openCloseDict = map[string]string{
	"(": ")",
	"[": "]",
	"{": "}",
	"<": ">",
}

func isOpening(char string) bool {
	if char == "(" || char == "[" || char == "{" || char == "<" {
		return true
	}
	return false
}

type stack struct {
	s []string
}

func makeStack() stack {
	var arr []string
	return stack{
		s: arr,
	}
}

func (s *stack) push(char string) {
	s.s = append(s.s, char)
}

func (s *stack) pop() string {
	l := len(s.s)
	res := s.s[l-1]
	s.s = s.s[:l-1]
	return res
}

func CompletionScore(input string) int {
	lines := strings.Split(input, "\n")
	var scores []int
	for _, line := range lines {
		compStr := CompletionString(line)
		if compStr != "" {
			score := 0
			for _, c := range strings.Split(compStr, "") {
				score = 5*score + completionPointDict[c]
			}
			scores = append(scores, score)
		}
	}

	sort.Slice(scores, func(i, j int) bool {
		return scores[i] > scores[j]
	})
	return scores[len(scores)/2]
}

var completionPointDict = map[string]int{
	")": 1,
	"]": 2,
	"}": 3,
	">": 4,
}

func CompletionString(line string) string {
	stack := makeStack()
	for _, char := range strings.Split(line, "") {
		if isOpening(char) {
			stack.push(char)
		} else {
			poppedChar := stack.pop()
			if openCloseDict[poppedChar] != char {
				return ""
			}
		}
	}
	compStr := ""
	for i := len(stack.s) - 1; i >= 0; i-- {
		compStr = compStr + openCloseDict[stack.s[i]]
	}
	return compStr
}
