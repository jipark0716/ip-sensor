package main

import (
	"github.com/jipark0716/ip-sensor/argument"
	"github.com/jipark0716/ip-sensor/pcap"
	"log"
)

func main() {
	arguments, err := argument.Get()
	if err != nil {
		log.Fatalf("argument parse fail %#v", err)
	}

	ipFilter, err := pcap.NewIpFilter(*arguments.Endpoint)
	if err != nil {
		log.Fatalf("endpoint filter 입력 실패 %#v", err)
	}

	pcap.Pcap(
		*arguments.Device,
		pcap.BPFQuery{
			Ip:   *arguments.Endpoint,
			Port: *arguments.Port,
		},
		[]pcap.Filter{
			ipFilter,
			pcap.TcpPshFilter{},
		},
		[]pcap.Formatter{
			pcap.PortFormat{},
			pcap.StringBodyFormatter{Offset: 4},
		},
	)
}
