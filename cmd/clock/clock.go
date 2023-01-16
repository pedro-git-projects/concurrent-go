package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, string(time.Now().Format("15:04:05\n")))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

// this clock is sequential, that is, it will only handle one client at a time
func clock1() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Listening on port 8000\n")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		handleConn(conn)
	}
}

// this is a concurrent version of clock1
func clock2() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Listening on port 8000\n")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

/*
Exercise 8.1: Modify clock2 to accept a port number, and write a program, clockwall,
that acts as a client of several clock servers at once,
reading the times from each one and displaying the results in a table,
akin to the wall of clocks seen in some business offices.
If you have access to geographically distributed computers, run instances remotely;
otherwise run local instances on different ports with fake time zones.
*/

func clockwall(port string) {
	_, err := strconv.Atoi(port)
	if err != nil {
		log.Fatal(err)
	}

	listener, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Listening on port :%s\n", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}

}
