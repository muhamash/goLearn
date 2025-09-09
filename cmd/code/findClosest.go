package code

import (
	"math"
)

func FindClosest(x int, y int, z int) int {
    dist1 := int(math.Abs(float64(x-z)))
	dist2 := int(math.Abs(float64(z-y)))

	if dist1 < dist2 {
		return  1
	} else if dist2 < dist1 {
		return  2
	} else {
		return  0
	}
}