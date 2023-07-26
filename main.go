package main

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"log"
)

func main() {
	//nics, err := net.Interfaces()
	//if err != nil {
	//	panic(err)
	//}
	//log.Printf("%#v\n", nics)

	if handle, err := pcap.OpenLive("lo0", 1600, true, pcap.BlockForever); err != nil {
		panic(err)
	} else if err := handle.SetBPFFilter("tcp and port 80"); err != nil { // optional
		panic(err)
	} else {
		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
		for packet := range packetSource.Packets() {
			handlePacket(packet) // Do something with a packet here.
		}
	}

}

func handlePacket(packet gopacket.Packet) {
	log.Printf("%#v\n", packet)
}
