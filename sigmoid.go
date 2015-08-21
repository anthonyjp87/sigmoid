package sigmoid

import (
	"math"
)

//if /False, return sigmoid. If True, return derivative of Sigmoid.
func Sigmoid(u float64, b bool) (r float64) {
	if b == true {
		r = u * (1 - u)
		return
	}
	return 1 / (1 + math.Exp(-u))
}
