package pcap

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

type Formatter interface {
	Run(packet gopacket.Packet) string
}

type PortFormat struct{}

func (p PortFormat) Run(packet gopacket.Packet) string {
	tcp, ok := packet.LayerClass(layers.LayerTypeTCP).(*layers.TCP)
	if !ok {
		return ""
	}
	return fmt.Sprintf(" %d -> %d", tcp.SrcPort, tcp.DstPort)
}

type StringBodyFormatter struct {
	Offset int
}

func (p StringBodyFormatter) Run(packet gopacket.Packet) string {
	tcp, ok := packet.LayerClass(layers.LayerTypeTCP).(*layers.TCP)
	if !ok || len(tcp.Payload) <= p.Offset {
		return ""
	}
	return string(tcp.Payload[p.Offset:])
}
