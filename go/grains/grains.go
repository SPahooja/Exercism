package grains

import (
	"fmt"
	"math"
)

// Square comment
func Square(n int) (uint64, error) {
	//n = math.Pow(float64(2), float64(n-1))
	var ans float64
	var n1 float64 = float64(n - 1)
	ans = math.Pow(2, n1)
	var final uint64
	final = uint64(ans)
	if n < 1 || n > 64 {
		return final, fmt.Errorf("Not in rangee")
	} else {
		return final, nil
	}
}

//Total comment
func Total() uint64 {
	return 18446744073709551615
}
