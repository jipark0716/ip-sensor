package pcap

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"log"
)

func Pcap(device string, filters ...Filter) {
	handle, err := pcap.OpenLive(device, 1600, true, pcap.BlockForever)
	if err != nil {
		log.Fatalf("pcap connect fail %#v", err)
	}

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		for _, f := range filters {
			if f.Run(packet) == false {
				break
			}
		}
		handlePacket(packet)
	}
}

func handlePacket(packet gopacket.Packet) {
	log.Printf("%#v\n", packet)
}
