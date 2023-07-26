package pcap

import (
	"github.com/google/gopacket"
	"net"
)

type Filter interface {
	Run(packet gopacket.Packet) bool
}

type IpFilter struct {
	ips []gopacket.Endpoint
}

func NewIpFilter(domain string) (Filter, error) {
	endpoints, err := net.LookupIP(domain)
	if err != nil {
		return nil, err
	}

	ips := make([]gopacket.Endpoint, len(endpoints))
	for i, endpoint := range endpoints {
		ips[i] = gopacket.NewEndpoint(1, endpoint)
	}
	return IpFilter{
		ips: ips,
	}, nil
}

func (f IpFilter) Run(packet gopacket.Packet) bool {
	ipHeader := packet.NetworkLayer()
	ipv4header := ipHeader.NetworkFlow()
	src := ipv4header.Src()
	for _, ip := range f.ips {
		if ip == src {
			return true
		}
	}
	return false
}
