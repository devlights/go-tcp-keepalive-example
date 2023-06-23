package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	ln, err := net.Listen("tcp", ":12345")
	if err != nil {
		return fmt.Errorf("error: net.Listen (%w)", err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			return fmt.Errorf("error: ln.Accept (%w)", err)
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	time.Sleep(10 * time.Second)
}
