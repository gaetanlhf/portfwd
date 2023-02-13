package main

import (
	"log"
	"net"
)

func udpForward(forward ForwardStruct) {
	srcAddr, err := net.ResolveUDPAddr(forward.Protocol, forward.From)
	if err != nil {
		log.Printf("Error resolving source address: %v\n", err)
	}

	dstAddr, err := net.ResolveUDPAddr(forward.Protocol, forward.To)
	if err != nil {
		log.Printf("Error resolving destination address: %v\n", err)
	}

	src, err := net.DialUDP(forward.Protocol, nil, srcAddr)
	if err != nil {
		log.Printf("The connection failed: %v", err)
	}
	defer src.Close()

	dst, err := net.ListenUDP(forward.Protocol, dstAddr)
	if err != nil {
		log.Printf("The connection failed: %v", err)
	}
	defer dst.Close()

	for {
		buf := make([]byte, 1024)
		n, _, err := src.ReadFromUDP(buf)
		if err != nil {
			log.Printf("Error reading from UDP socket: %v\n", err)
		}

		_, err = dst.WriteToUDP(buf[:n], dstAddr)
		if err != nil {
			log.Printf("Error forwarding UDP packet: %v\n", err)
		}
	}
}
