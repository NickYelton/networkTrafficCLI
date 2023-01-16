package main

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"log"
)

func main() {
	// Open a live connection to the network interface "en0 <-- change to your target network"
	handle, err := pcap.OpenLive("en0", 1600, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Set filter to capture only TCP packets with destination port 80 and source port range of 3000-3010
	filter := "tcp and dst port 80 and src portrange 3000-3010"
	err = handle.SetBPFFilter(filter)
	if err != nil {
		log.Fatal(err)
	}

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		// Extract the TCP layer from the packet
		if tcpLayer := packet.Layer(layers.LayerTypeTCP); tcpLayer != nil {
			tcp, _ := tcpLayer.(*layers.TCP)
			log.Printf("Source Port: %d, Destination Port: %d", tcp.SrcPort, tcp.DstPort)
		}
		// Extract the IP layer from the packet
		if ipLayer := packet.NetworkLayer(); ipLayer != nil {
			switch ip := ipLayer.(type) {
			case *layers.IPv4:
				log.Printf("Protocol: %d", ip.Protocol)
			case *layers.IPv6:
				log.Printf("Protocol: %d", ip.NextHeader)
			}
		}

	}
}
