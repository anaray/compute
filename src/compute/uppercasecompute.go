package compute

import (
	"strings"
)

type UpperCasePlugin struct{}

func UpperCase() *UpperCasePlugin {
	return new(UpperCasePlugin)
}

func (p *UpperCasePlugin) Execute(arg Args) {
	for {
		message := <-arg.Incoming
		arg.Outgoing <- strings.ToUpper(message)
	}
}
