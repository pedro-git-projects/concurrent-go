package main

import (
	"fmt"

	"github.com/pedro-git-projects/concurrent-go/pkg"
)

func main() {
	rng := pkg.RNG()
	for i := 0; i < 10; i++ {
		r := rng.Intn(10)
		fmt.Println(r)
	}
}
