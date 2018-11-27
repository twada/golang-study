package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

var p = flag.Int("p", 8000, "port number")

func main() {
	flag.Parse()
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *p))
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // クライアントとの接続が切れた
		}
		time.Sleep(1 * time.Second)
	}
}
