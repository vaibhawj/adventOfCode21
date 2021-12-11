package day8

import (
	"strconv"
	"strings"
)

// Seven segment digits

func EasyDecode(input string) int {

	entries := strings.Split(input, "\n")
	count := 0
	for _, entry := range entries {
		output := strings.Split(entry, "|")[1]
		for _, o := range strings.Split(output, " ") {
			length := len(o)
			if length == 2 || length == 3 || length == 4 || length == 7 {
				count++
			}
		}
	}
	return count
}

func RealDecode(input string) int {
	entries := strings.Split(input, "\n")
	sum := 0
	for _, entry := range entries {
		sum += Decode(entry)
	}
	return sum
}

func Decode(entry string) int {
	input := strings.Split(strings.Trim(strings.Split(entry, "|")[0], " "), " ")
	output := strings.Split(strings.Trim(strings.Split(entry, "|")[1], " "), " ")
	knownStrToNumber := make(map[string]int)
	knownNumberToStr := make(map[int]string)
	lengthStringsMap := make(map[int][]string)

	for _, s := range input {
		length := len(s)
		lengthStringsMap[length] = append(lengthStringsMap[length], s)
	}
	knownStrToNumber[lengthStringsMap[2][0]] = 1
	knownNumberToStr[1] = lengthStringsMap[2][0]

	knownStrToNumber[lengthStringsMap[3][0]] = 7
	knownNumberToStr[7] = lengthStringsMap[3][0]

	knownStrToNumber[lengthStringsMap[4][0]] = 4
	knownNumberToStr[4] = lengthStringsMap[4][0]

	knownStrToNumber[lengthStringsMap[7][0]] = 8
	knownNumberToStr[8] = lengthStringsMap[7][0]

	// 3
	for _, s := range lengthStringsMap[5] {
		one := strings.Split(lengthStringsMap[2][0], "")
		if strings.Contains(s, one[0]) && strings.Contains(s, one[1]) {
			knownStrToNumber[s] = 3
			knownNumberToStr[3] = s
			break
		}
	}

	// 6
	for _, s := range lengthStringsMap[6] {
		one := strings.Split(lengthStringsMap[2][0], "")
		if !strings.Contains(s, one[0]) || !strings.Contains(s, one[1]) {
			knownStrToNumber[s] = 6
			knownNumberToStr[6] = s
			break
		}
	}

	for _, s := range lengthStringsMap[6] {
		leftSideChars := strings.Split(Minus(knownNumberToStr[8], knownNumberToStr[3]), "")
		if s == knownNumberToStr[6] {
			continue
		} else {
			if strings.Contains(s, leftSideChars[0]) && strings.Contains(s, leftSideChars[1]) {
				knownStrToNumber[s] = 0
				knownNumberToStr[0] = s
			} else {
				knownStrToNumber[s] = 9
				knownNumberToStr[9] = s
			}
		}
	}

	for _, s := range lengthStringsMap[5] {
		if s == knownNumberToStr[3] {
			continue
		} else {
			if len(Minus(s, knownNumberToStr[9])) == 0 {
				knownStrToNumber[s] = 5
				knownNumberToStr[5] = s
			} else {
				knownStrToNumber[s] = 2
				knownNumberToStr[2] = s
			}
		}
	}

	decodedOutputStr := ""
	for _, o := range output {
		for k, v := range knownStrToNumber {
			if len(o) == len(k) && len(Minus(k, o)) == 0 && len(Minus(o, k)) == 0 {
				decodedOutputStr += strconv.Itoa(v)
			}
		}
	}
	decodedOutput, _ := strconv.Atoi(decodedOutputStr)

	return decodedOutput
}

func Minus(str1, str2 string) string {
	result := ""
	for _, s := range strings.Split(str1, "") {
		if !strings.Contains(str2, s) {
			result = result + s
		}
	}
	return result
}
