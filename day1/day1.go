package day1

func Part1(inputs []int) int {
	var counter = 0
	for i := 1; i < len(inputs); i++ {
		if inputs[i] > inputs[i-1] {
			counter++
		}
	}
	return counter
}

func Part2(inputs []int) int {
	var counter = 0
	for i := 0; i < len(inputs); i++ {
		if i+3 == len(inputs) {
			break
		}
		if inputs[i+1]+inputs[i+2]+inputs[i+3] > inputs[i]+inputs[i+1]+inputs[i+2] {
			counter++
		}
	}
	return counter
}
