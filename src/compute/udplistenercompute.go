package compute

import (
	"log"
	"net"
)

type UDPListenerPlugin struct{}

func UDPListener() *UDPListenerPlugin {
	return new(UDPListenerPlugin)
}

func (p *UDPListenerPlugin) Execute(arg Args) {
	var udpaddr *net.UDPAddr
	var err error
	if udpaddr, err = net.ResolveUDPAddr("udp4", "127.0.0.1:1200"); err != nil {
		return
	}

	conn, err := net.ListenUDP("udp4", udpaddr)
	for {
		buffer := make([]byte, 1024)
		if c, addr, err := conn.ReadFromUDP(buffer); err != nil {
			log.Printf("error: %d byte datagram from %s with error %s\n", c, addr.String(), err.Error())
			return

		} else {
			arg.Outgoing <- string(buffer[:c])
		}
	}
}
