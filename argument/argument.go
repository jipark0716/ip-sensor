package argument

import (
	"flag"
	"fmt"
	"github.com/charmbracelet/huh"
	"net"
)

type Arguments struct {
	Device   *string
	Endpoint *string
	Port     *string
}

func Get() (result Arguments, err error) {
	result.Device = flag.String("device", "", "a string")
	result.Port = flag.String("port", "", "a string")
	result.Endpoint = flag.String("endpoint", "", "a string")

	flag.Parse()

	if result.Device == nil || *result.Device == "" {
		err = getNetworkInterface(result.Device)
		if err != nil {
			return
		}
	}

	if result.Endpoint == nil || *result.Endpoint == "" {
		err = getEndpoint(result.Endpoint)
		if err != nil {
			return
		}
	}

	if result.Port == nil || *result.Port == "" {
		err = getPort(result.Port)
		if err != nil {
			return
		}
	}

	return
}

func getEndpoint(result *string) error {
	return huh.NewInput().
		Title("filter endpoint").
		Value(result).
		Run()
}

func getPort(result *string) error {
	return huh.NewInput().
		Title("filter port").
		Value(result).
		Run()
}

func getNetworkInterface(result *string) error {
	interfaces, err := net.Interfaces()
	if err != nil {
		return err
	}

	var options = make([]huh.Option[string], len(interfaces))
	for index, element := range interfaces {
		options[index] = huh.NewOption(fmt.Sprintf("mtu:%d\tname:%s", element.MTU, element.Name), element.Name)
	}

	return huh.NewSelect[string]().
		Title("Choose network interface").
		Options(options...).
		Value(result).
		Run()
}
