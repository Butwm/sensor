package main

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"log"
)

func main() {
	iface := "eth0" // Interface
	targetIP := "178.143.37.189" // Your ip adress

	handle, err := pcap.OpenLive(iface, 1600, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	packetCount := 0

	for packet := range packetSource.Packets() {
		networkLayer := packet.NetworkLayer()
		if networkLayer != nil {
			dstIP := networkLayer.NetworkFlow().Dst().String()

			if dstIP == targetIP {
				packetCount++
			}
		}

		if packetCount%10 == 0 {
			fmt.Printf("Počet paketov pre %s: %d\n", targetIP, packetCount)
		}
	}
}
