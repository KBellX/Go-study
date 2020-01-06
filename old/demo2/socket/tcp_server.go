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

		go handleClient(conn)

	}

	time.Sleep(180e9)
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	//  阻塞A请求，不会影响B请求, 体现为多次请求同时返回。
	time.Sleep(1e9)
	_, _ = conn.Write([]byte("adwad"))
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
