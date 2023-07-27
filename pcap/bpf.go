package pcap

import (
	"fmt"
)

type BPFQuery struct {
	Ip   string
	Port string
}

func (q BPFQuery) String() string {
	return fmt.Sprintf("tcp and port %s", q.Port)
}
