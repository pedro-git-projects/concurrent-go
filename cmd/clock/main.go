package main

import (
	"flag"
	"fmt"
)

func main() {
	port := flag.String("port", "8000", "connection port")
	flag.Parse()
	fmt.Println(*port)
	clockwall(*port)
}
