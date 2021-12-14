package day6

func NumberOfLanternFish(fishStates []int, days int) int {
	current := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	for _, f := range fishStates {
		current[f]++
	}

	for d := 1; d <= days; d++ {
		nextGen := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
		copy(nextGen, current)

		for i := 0; i <= 8; i++ {
			if i == 8 {
				if current[8] > 0 {
					nextGen[8]--
				}
				continue
			}
			nextGen[i] = current[i+1]
		}
		nextGen[8] = current[0]
		nextGen[6] += current[0]

		current = nextGen
		//fmt.Println(current)
	}

	sum := 0
	for _, f := range current {
		sum += f
	}

	return sum
}
