package main

import (
	_ "embed"
	"fmt"
	"math"
)

func charge_time(race_time, distance int) (float64, float64) {
	sol1 := 0.5*float64(race_time) + math.Sqrt(math.Pow(float64(-race_time)/2, 2)-float64(distance))
	sol2 := 0.5*float64(race_time) - math.Sqrt(math.Pow(float64(-race_time)/2, 2)-float64(distance))
	return sol1, sol2
}

func ints_between(a, b float64) []int {
	var ints []int

	if a > b {
		a, b = b, a
	}
	smaller := int(math.Ceil(a))
	bigger := int(math.Floor(b))
	fmt.Println(a, b, smaller, bigger)
	for i := smaller; i <= bigger; i++ {
		if float64(i) == a || float64(i) == b {
			continue
		}
		ints = append(ints, i)
	}
	return ints
}

func main() {

	// example:
	// ans1 := ints_between(charge_time(7, 9))
	// ans2 := ints_between(charge_time(15, 40))
	// ans3 := ints_between(charge_time(30, 200))
	// fmt.Println(len(ans1) * len(ans2) * len(ans3))

	// input part_a
	ans1 := ints_between(charge_time(40, 219))
	ans2 := ints_between(charge_time(81, 1012))
	ans3 := ints_between(charge_time(77, 1365))
	ans4 := ints_between(charge_time(72, 1089))
	fmt.Println(ans1, ans2, ans3, ans4)
	fmt.Println(len(ans1), len(ans2), len(ans3), len(ans4))
	fmt.Println(len(ans1) * len(ans2) * len(ans3) * len(ans4))

}
