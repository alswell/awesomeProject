package main

import (
	"bufio"
	"fmt"
	"metric_monitor/common/pcrypto"
	"net"
)

func ProcessRequest(conn net.Conn) {
	defer conn.Close()

	pc := pcrypto.Pcrypto{}
	pc.Init([]byte{0x1f, 0x32, 0xb0, 0x29, 0xd7, 0x3f, 0xba, 0x56, 0x22, 0x71, 0xa8, 0x4d, 0x52, 0xe9, 0xa2, 0x3f})
	reader := bufio.NewReader(conn)
	buff := make([]byte, 1024)
	for {
		n, e := reader.Read(buff)
		fmt.Println(n, e)
		if e != nil {
			break
		}
		a, err := pc.Decompression(buff)
		fmt.Println(string(a), err)
	}
}

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "0.0.0.0:3008")
	fmt.Println(tcpAddr, err)
	l, err := net.ListenTCP("tcp", tcpAddr)
	fmt.Println(l, err)
	defer l.Close()

	fmt.Printf("server started %v\n", tcpAddr)
	for {
		conn, err := l.AcceptTCP()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go ProcessRequest(conn)
	}
}
