package wardrobe

import (
	"testing"
)

var wardrobeElements []int
var spaceLength int
var c [][]int

func init() {
	wardrobeElements = []int{50, 75, 100, 120}
	spaceLength = 250
	c = combinations(wardrobeElements, spaceLength)
}

func TestFiftyFit(t *testing.T) {
	if expectList := []int{50, 50, 50, 50, 50}; !listContains(expectList, c) {
		t.Error("expect 50, 50 ,50 ,50 ,50 not found")
	}
}

func TestFit2(t *testing.T) {
	if expectList := []int{50, 100, 100}; !listContains(expectList, c) {
		t.Error("expect 50, 100, 100 not found")
	}
}

func TestCalcMoney(t *testing.T) {
	var l []int = []int{50, 75, 100, 120, 50}
	if 381 != calcMoney(l) {
		t.Error("expect 318, actual:", calcMoney(l))
	}
}

func TestCheapestOne(t *testing.T) {
	var l [][]int
	l = append(l, []int{50, 75, 100, 120, 50})
	l = append(l, []int{50, 75, 100})

	var cheapest []int
	var money int

	cheapest, money = cheapestOne(l)

	if( !listEquals(l[1], cheapest)){
		t.Error("cheapest list is wrong")
	}

	if( calcMoney(l[1]) != money){
		t.Error("expect money is wrong, expect:", calcMoney(l[1]), "actual:", money)
	}

}

func listContains(target []int, list [][]int) bool {
	for _, l := range list {
		if listEquals(target, l) {
			return true
		}
	}

	return false
}

func listEquals(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
