package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	listener, e := net.Listen("tcp", "127.0.0.1:8000")
	if e != nil {
		log.Fatal(e)
	}
	for {
		conn, e := listener.Accept()
		if e != nil {
			log.Println(e)
			continue
		}
		go handleConnc(conn)

	}

}

func handleConnc(conn net.Conn) {
	defer conn.Close()

	ipAddr := conn.RemoteAddr().String()
	bytes := make([]byte, 1024)

	for {
		n, err := conn.Read(bytes)
		if err != nil {
			fmt.Println(err)
			return
		}
		//切片
		result := bytes[:n]
		fmt.Printf("接收到数据来自[%s]==>[%d]:%s\n", ipAddr, n, string(result))
		if "exit" == string(result) {
			fmt.Println(ipAddr, "退出连接")
			return
		}
		conn.Write([]byte(strings.ToUpper(string(result))))

	}

}
