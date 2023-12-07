package main

import (
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
	for i := smaller; i <= bigger; i++ {
		if float64(i) == a || float64(i) == b {
			continue
		}
		ints = append(ints, i)
	}
	return ints
}

func main() {
	ans1 := ints_between(charge_time(40817772, 219101213651089))
	fmt.Println(len(ans1))
}
