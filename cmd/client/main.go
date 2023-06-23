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
	conn, err := net.Dial("tcp", ":12345")
	if err != nil {
		return fmt.Errorf("error: net.Dial (%w)", err)
	}
	defer conn.Close()

	tcpConn, ok := conn.(*net.TCPConn)
	if !ok {
		return fmt.Errorf("error: not *net.TCPConn")
	}

	// Set Keep-Alive
	{
		if err := tcpConn.SetKeepAlive(true); err != nil {
			return fmt.Errorf("error: tcpConn.SetKeepAlive (%w)", err)
		}

		if err := tcpConn.SetKeepAlivePeriod(2 * time.Second); err != nil {
			return fmt.Errorf("error: tcpConn.SetKeepAlivePeriod (%w)", err)
		}
	}

	time.Sleep(10 * time.Second)

	return nil
}
