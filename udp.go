package main

import (
	"log"
	"net"

	"github.com/libp2p/go-reuseport"
)

func udpForward(forward ForwardStruct) {

	dstAddr, err := net.ResolveUDPAddr(forward.Protocol, forward.To)
	if err != nil {
		log.Printf("Error resolving destination address: %v\n", err)
	}

	src, err := reuseport.ListenPacket(forward.Protocol, forward.From)
	if err != nil {
		log.Printf("The connection failed: %v", err)
	}
	defer src.Close()

	dst, err := net.DialUDP(forward.Protocol, nil, dstAddr)
	if err != nil {
		log.Printf("The connection failed: %v", err)
	}
	defer dst.Close()

	for {
		buf := make([]byte, 1600)
		n, _, err := src.ReadFrom(buf)
		if err != nil {
			log.Printf("Error reading from UDP socket: %v\n", err)
		}

		_, _ = dst.Write(buf[:n])

	}
}
