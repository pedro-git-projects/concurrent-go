package main

import (
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	l := new(sync.Locker)
	sleeping(10, wg)
	mutexIncDec()
	mutexRW(wg, *l)
}
