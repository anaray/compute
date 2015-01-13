package compute

import (
	"os"
	"sync"
	//"github.com/BurntSushi/toml"
)

type Args struct {
	Incoming chan Packet
	Outgoing chan Packet
	WaitGroup *sync.WaitGroup
	Container *map[string]interface{}
	Logger    *Log
}

type Computes interface {
	String() string
	Execute(*Args)
}

type Packet map[string]interface{}

func Run(computes ...Computes) {
	in := make(chan Packet, 10000)
	logger := Logger(os.Stdout)
	var wg sync.WaitGroup
	var arg Args
	var args  []Args
	logger.Logf("Initializing Compute ...")
	for i := 0; i < len(computes); i++ {
		out := make(chan Packet, 10000)
		arg = Args{Incoming: in, Outgoing: out, Logger: logger, WaitGroup: &wg}
		args = append(args, arg)
		in = out
	}
	for indx, compute := range computes {
		arg.WaitGroup.Add(1)
		logger.Logf("Initializing Compute: %s", compute.String())
		go compute.Execute(&args[indx])
	}
	arg.WaitGroup.Wait()
}

func NewPacket() Packet {
	packet := make(Packet)
	return packet
}
