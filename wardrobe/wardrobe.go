package wardrobe

func combinations(wardrobeElements []int, spaceLength int) [][]int {
	var c [][]int
	c = augment(wardrobeElements, spaceLength, c)

	var result [][]int

	result = threeVar(c, spaceLength, result)
	result = fourVar(c, spaceLength, result)
	result = fiveVar(c, spaceLength, result)

	return result
}

func calcMoney(l []int) int {
	var money int
	for _, i := range l {
		switch i {
		case 50:
			money += 59
		case 75:
			money += 62
		case 100:
			money += 90
		case 120:
			money += 111
		}
	}

	return money
}

func cheapestOne(l [][]int) ([]int, int) {
	var money int
	var index int
	for i, v := range l {
		if m := calcMoney(v); money == 0 || m < money {
			money = m
			index = i
		}
	}

	return l[index], money
}

func fiveVar(c [][]int, spaceLength int, result [][]int) [][]int {
	for i := 0; i < 4; i++ {
		for i1 := i; i1 < 4; i1++ {
			for i2 := i1; i2 < 4; i2++ {
				for i3 := i2; i3 < 4; i3++ {
					for i4 := i3; i4 < 4; i4++ {
						if (c[0][i] + c[1][i1] + c[2][i2] + c[3][i3] + c[4][i4]) == spaceLength {
							result = append(result, []int{c[0][i], c[1][i1], c[2][i2], c[3][i3], c[4][i4]})
						}
					}
				}
			}
		}
	}
	return result
}

func fourVar(c [][]int, spaceLength int, result [][]int) [][]int {
	for i := 0; i < 4; i++ {
		for i1 := i; i1 < 4; i1++ {
			for i2 := i1; i2 < 4; i2++ {
				for i3 := i2; i3 < 4; i3++ {
					if (c[0][i] + c[1][i1] + c[2][i2] + c[3][i3]) == spaceLength {
						result = append(result, []int{c[0][i], c[1][i1], c[2][i2], c[3][i3]})
					}
				}
			}
		}
	}
	return result
}

func threeVar(c [][]int, spaceLength int, result [][]int) [][]int {
	for i := 0; i < 4; i++ {
		for i1 := i; i1 < 4; i1++ {
			for i2 := i1; i2 < 4; i2++ {
				if (c[0][i] + c[1][i1] + c[2][i2]) == spaceLength {
					result = append(result, []int{c[0][i], c[1][i1], c[2][i2]})
				}
			}
		}
	}
	return result
}

func augment(wardrobeElements []int, spaceLength int, c [][]int) [][]int {
	for i := 0; i < dimensions(wardrobeElements, spaceLength); i++ {
		c = append(c, wardrobeElements)
	}
	return c
}

func dimensions(wardrobeElements []int, spaceLength int) int {
	return spaceLength / wardrobeElements[0]
}
