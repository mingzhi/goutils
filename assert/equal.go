package assert

import (
	"math"
)

// Asserts that two floats are equal to within a positive delta.
// typ = 0: Absolute delta; typ = 1: Relative delta.
// NaNs or Infs are considerred equal.
func EqualFloat64(actual, expected, delta float64, typ int) (status bool) {
	switch {
	case math.IsNaN(actual) || math.IsNaN(expected):
		status = math.IsNaN(actual) == math.IsNaN(expected)
		break
	case math.IsInf(actual, 0) || math.IsInf(expected, 0):
		status = math.IsInf(actual, 0) == math.IsInf(expected, 0)
		break
	case expected == 0:
		status = math.Abs(actual-expected) < math.Abs(delta)
		break
	case expected != 0:
		if typ == 0 {
			status = math.Abs(actual-expected) < math.Abs(delta)
		} else {
			status = math.Abs(actual-expected)/math.Abs(expected) < math.Abs(delta)
		}
	}
	return
}
