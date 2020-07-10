// Copyright © 2020 Bin Liu <bin.liu@enmotech.com>

package utils

import (
	"math"
)

// CalcFloat64 round
func CalcFloat64(f float64, pos int) float64 {
	p := math.Pow10(pos)
	a := f * p
	// a = a + 0/5
	return math.Round(a) / p
}
