package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

const (
	NETWORK = "tcp"
	TCPADDR = "127.0.0.1:9800"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr(NETWORK, TCPADDR)
	checkError(err)

	conn, err := net.DialTCP(NETWORK, nil, tcpAddr)
	checkError(err)
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)
	result, err := ioutil.ReadAll(conn)
	checkError(err)
	fmt.Println(string(result))
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
