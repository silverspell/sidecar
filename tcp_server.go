package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
)

func createServer() error {
	listener, err := net.Listen(TYPE, HOST+PORT)
	if err != nil {
		return err
	}

	defer listener.Close()
	fmt.Println("Listening!")
	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		go handleIncoming(conn)
	}
}

func handleIncoming(conn net.Conn) {
	reader := bufio.NewReader(conn)
	buffer := bytes.Buffer{}

	for {
		n, isPrefix, err := reader.ReadLine()
		if err != nil {
			fmt.Printf("%+v\n", err)
			return
		}

		buffer.Write(n)

		if buffer.Len() > 0 && !isPrefix {
			fmt.Printf("%+v\n", buffer.Bytes())
			conn.Write(buffer.Bytes())
			conn.Write([]byte{'\n'})
			buffer.Reset()
		}
	}
}
