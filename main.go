package main

import (
	"fmt"
	"github.com/charmbracelet/huh"
	"github.com/jipark0716/ip-sensor/pcap"
	"log"
	"net"
)

func main() {
	device, err := selectNetworkInterface()
	if err != nil {
		log.Fatalf("network interface 선택 실패 %#v", err)
	}

	var endpoint string
	err = huh.NewInput().
		Title("filter endpoint").
		Value(&endpoint).
		Run()

	filter, err := pcap.NewIpFilter(endpoint)

	if err != nil {
		log.Fatalf("ip filter 입력 실패 %#v", err)
	}

	pcap.Pcap(device, filter)
}

func selectNetworkInterface() (result string, err error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return
	}

	var options = make([]huh.Option[string], len(interfaces))
	for index, element := range interfaces {
		options[index] = huh.NewOption(fmt.Sprintf("mtu:%d\tname:%s", element.MTU, element.Name), element.Name)
	}

	err = huh.NewSelect[string]().
		Title("Choose network interface").
		Options(options...).
		Value(&result).
		Run()

	return
}
