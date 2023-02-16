package main

import (
	"io"
	"log"
	"net"

	"github.com/libp2p/go-reuseport"
)

func tcpForward(forward ForwardStruct) {
	listener, err := reuseport.Listen(forward.Protocol, forward.To)

	if err != nil {
		log.Printf("The connection failed: %v", err)
	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Printf("The connection was not accepted: %v", err)
		}

		client, err := net.Dial(forward.Protocol, forward.From)

		if err != nil {
			log.Printf("The connection failed: %v", err)
			defer conn.Close()
			return
		}

		go func() {
			defer client.Close()
			defer conn.Close()
			io.Copy(client, conn)
		}()

		go func() {
			defer client.Close()
			defer conn.Close()
			io.Copy(conn, client)
		}()
	}
}
