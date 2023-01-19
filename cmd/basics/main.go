package main

import (
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	sleeping(10, wg)
	commonIncrement()
}
