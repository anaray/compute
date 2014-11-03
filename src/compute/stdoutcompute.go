package compute

import (
	"fmt"
)

type StdoutPlugin struct{}

func Stdout() *StdoutPlugin {
	return new(StdoutPlugin)
}

func (p *StdoutPlugin) Execute(arg Args) {
	for {
		message := <-arg.Incoming
		fmt.Println("word :" + message)
	}
}
