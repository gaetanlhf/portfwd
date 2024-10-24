package main

import (
	"log"
	"net"

	"github.com/libp2p/go-reuseport"
)

func udpForward(forward ForwardStruct) {
	src, err := reuseport.ListenPacket(forward.Protocol, forward.From)
	if err != nil {
		log.Printf("The connection failed: %v", err)
	}
	defer src.Close()

	var sliceDst []*net.UDPConn

	for _, to := range forward.To {
		dstAddr, err := net.ResolveUDPAddr(forward.Protocol, to)
		if err != nil {
			log.Printf("Error resolving destination address: %v\n", err)
		}

		dst, err := net.DialUDP(forward.Protocol, nil, dstAddr)
		if err != nil {
			log.Printf("The connection failed: %v", err)
		}
		dst.Close()

		sliceDst = append(sliceDst, dst)
	}

	for {
		buf := make([]byte, 65535)
		n, _, err := src.ReadFrom(buf)
		if err != nil {
			log.Printf("Error reading from UDP socket: %v\n", err)
		}

		for _, dst := range sliceDst {
			_, _ = dst.Write(buf[:n])
		}
	}
}
