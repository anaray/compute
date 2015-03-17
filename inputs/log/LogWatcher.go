package main

import (
	"fmt"
	"github.com/ActiveState/tail"
	"github.com/anaray/compute"
	"github.com/anaray/regnet"
	"os"
	"path/filepath"
	"regexp"
)

type LogWatcher struct {
}

func NewLogWatcher() *LogWatcher {
	return &LogWatcher{}
}

func (lw *LogWatcher) string() string {
	return "LogWatcher"
}

func (lw *LogWatcher) Execute(arg *compute.Args) {
	r, _ := regexp.Compile("client_(.*).log")
	//folder := "/home/suresh/official/arb/poc/MetricStream_EGRCP/SYSTEMi/Systemi/log/"
	folder <- arg.Incoming

	nc := make(chan string)

	go func() {
		for event := range nc {
			fmt.Println("tailing file ", event)
			if r.MatchString(filepath.Base(event)) {
				Watch(event)
			}
		}
	}()
	notifier := NewNotifier(nc, folder)
	notifier.Watch()

}

func Watch(file string) {
	go func() {
		var lg *Logs
		for {
			t, _ := tail.TailFile(file, tail.Config{Follow: true, Location: &tail.SeekInfo{0, os.SEEK_END}})
			for line := range t.Lines {
				exists, _ := regnet.Exists(line, "%{MS_DELIM}")
				if exists {
					if lg != nil && len(lg.Store) > 0 {
						//push it further
						packet := compute.NewPacket()
						packet["LOG"] = lg
						packet["SOURCE"] = file
						arg.Outgoing <- packet
					}
					lg = NewLog(line)
				} else {
					if len(lg.Store) > 0 {
						store := []byte(lg.Store)
						lg.Store = string(append(store[:], line[:]...))
					}
				}
			}

		}
	}()
}

func NewLog(line []byte) *Logs {
	log := new(Logs)
	log.Store = string(line)
	return log
}

type Logs struct {
	Store string
}

func main() {
	var input string
	fmt.Scanln(&input)
	fmt.Println("Exit")
}
