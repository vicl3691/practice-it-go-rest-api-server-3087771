package mytesting

import (
	"testing"
	"math"
)

func almostEqual(v1, v2 float64) bool {
	return math.Abs(v1-v2) <= 0.001
}

func TestSimple (t *testing.T) {
	val, err := Sqrt(2)

	if err != nil {
		t.Fatalf("error in calculations - %s", err)
	}

	if !almostEqual(val, 1.414214) {
		t.Fatalf("bad value - %f", val)
	}
}

