package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strconv"
	"time"
)

const (
	NETWORK = "tcp"
	TCPADDR = "127.0.0.1:9800"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr(NETWORK, TCPADDR)
	checkError(err)

	num := 3
	for i := 1; i <= num; i++ {
		go doConn(tcpAddr, i)
	}

	time.Sleep(3 * time.Second)
}

func doConn(tcpAddr *net.TCPAddr, i int) {
	conn, err := net.DialTCP(NETWORK, nil, tcpAddr)
	checkError(err)
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)
	result, err := ioutil.ReadAll(conn)
	checkError(err)
	fmt.Println(strconv.Itoa(i) + ":" + string(result))
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		// os.Exit(1)
	}
}
