package assert

import "testing"
import "math"

func TestEqualFloat6464(t *testing.T) {
	var actual, expected, delta float64
	var assert bool
	// One is NaN
	actual = math.NaN()
	expected = 0.1
	assert = EqualFloat64(actual, expected, delta, 0)
	if assert != false {
		t.Errorf("actual: %f, expected: %f\n", actual, expected)
	}

	// Both are NaN
	actual = math.NaN()
	expected = math.NaN()
	assert = EqualFloat64(actual, expected, delta, 0)
	if assert != true {
		t.Errorf("actual: %f, expected: %f\n", actual, expected)
	}

	// One is Inf
	actual = math.Inf(0)
	expected = 0.001
	assert = EqualFloat64(actual, expected, delta, 0)
	if assert != false {
		t.Errorf("actual: %f, expected: %f\n", actual, expected)
	}

	// Both are Inf
	actual = math.Inf(0)
	expected = math.Inf(0)
	assert = EqualFloat64(actual, expected, delta, 0)
	if assert != true {
		t.Errorf("actual: %f, expected: %f\n", actual, expected)
	}

	// Expected is 0, false case
	actual = 0.23
	expected = 0
	delta = 0.01
	assert = EqualFloat64(actual, expected, delta, 0)
	if assert != false {
		t.Errorf("actual: %f, expected: %f, delta: %f\n", actual, expected, delta)
	}

	// Expected is 0, true case
	actual = 0.003
	expected = 0
	delta = 0.01
	assert = EqualFloat64(actual, expected, delta, 0)
	if assert != true {
		t.Errorf("actual: %f, expected: %f, delta: %f\n", actual, expected, delta)
	}

	// Expected is not 0, true case, with Absolute delta
	actual = 0.209
	expected = 0.2
	delta = 0.01
	assert = EqualFloat64(actual, expected, delta, 0)
	if assert != true {
		t.Errorf("actual: %f, expected: %f, delta: %f\n", actual, expected, delta)
	}

	// Expected is not 0, true case, with Relative delta
	actual = 0.23
	expected = 0.2
	delta = 0.01
	assert = EqualFloat64(actual, expected, delta, 0)
	if assert != false {
		t.Errorf("actual: %f, expected: %f, delta: %f\n", actual, expected, delta)
	}

	// Expected is not 0, true case, with Relative delta
	actual = 0.0000034444444
	expected = 0.000003444445
	delta = 0.001
	assert = EqualFloat64(actual, expected, delta, 1)
	if assert != true {
		t.Errorf("actual: %f, expected: %f, delta: %f\n", actual, expected, delta)
	}
}
