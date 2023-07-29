package pcap

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"log"
	"strings"
)

func Pcap(device string, query BPFQuery, filters []Filter, formatters []Formatter) {
	handle, err := pcap.OpenLive(device, 1600, true, pcap.BlockForever)
	if err != nil {
		log.Fatalf("pcap connect fail %#v", err)
	}

	err = handle.SetBPFFilter(query.String())
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
		logs := make([]string, len(formatters))
		for i, formatter := range formatters {
			logs[i] = formatter.Run(packet)
		}
		log.Println(strings.Join(logs, " "))
	}
}
