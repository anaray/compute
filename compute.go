package compute

import (
	"os"
	"sync"
	"runtime"
	//"github.com/BurntSushi/toml"
)

type Args struct {
	Incoming  chan Packet
	Outgoing  chan Packet
	WaitGroup *sync.WaitGroup
	Store     *Cache
	Logger    *Log
}

type Computes interface {
	String() string
	Execute(*Args)
}

type Packet map[string]interface{}

func Run(computes ...Computes) {
	ncpu := runtime.NumCPU()
	runtime.GOMAXPROCS(ncpu)
	in := make(chan Packet, 100)
	logger := Logger(os.Stdout)
	var wg sync.WaitGroup
	var arg Args
	var args []Args
	cache := NewCache()

	logger.Logf("Initializing Compute ...")
	for i := 0; i < len(computes); i++ {
		out := make(chan Packet, 100)
		arg = Args{Incoming: in, Outgoing: out, WaitGroup: &wg, Store: cache, Logger: logger}
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

type Cache struct {
	items map[string]interface{}
	lock  *sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		items: make(map[string]interface{}, 1024),
		lock:  new(sync.RWMutex),
	}
}

func (c *Cache) Add(key string, item interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.items[key] = item
}

func (c *Cache) Get(key string) interface{} {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.items[key]
}
