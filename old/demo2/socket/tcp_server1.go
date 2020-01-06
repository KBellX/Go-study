package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

const (
	NETWORK = "tcp"
	TCPADDR = ":9800"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr(NETWORK, TCPADDR)
	checkError(err)
	listener, err := net.ListenTCP(NETWORK, tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		time.Sleep(1 * time.Second)
		// params, err := conn.Read()
		_, err = conn.Write([]byte("adwad"))
		if err != nil {
			continue
		}
		conn.Close()
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
