package main

import (
	"math/rand"
	"time"
)

func getRandomNumber(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max) + 1
}
