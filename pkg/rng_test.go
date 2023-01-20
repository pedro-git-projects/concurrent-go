package pkg

import (
	"testing"
)

func TestRNG(t *testing.T) {
	rng := RNG()
	acc := make([]int, 100)
	for i := 0; i < 100; i++ {
		r := rng.Intn(100)
		if Contains(acc, r) {
		} else {
			acc = append(acc, r)
		}
	}
	if len(acc) == 1 {
		t.Error("Random number generator is (probably) not working")
	}
}
