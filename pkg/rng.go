package pkg

import (
	"math/rand"
	"time"
)

func RNG() *rand.Rand {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r
}
