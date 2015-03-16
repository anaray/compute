package compute

import (
	"testing"
	"strings"
)

var input = "value"
var result string

/** Compute : CreatorCompute definition start **/
type CreatorCompute struct{}

func getCreatorCompute() *CreatorCompute {
	return &CreatorCompute{}
}

func (c1 *CreatorCompute) String() string {
	return "CreatorCompute"
}

func (c1 *CreatorCompute) Execute(arg *Args) {
	defer arg.WaitGroup.Done()
	packet := NewPacket()
	packet["key"] = "value"
	arg.Outgoing <- packet
}
/** Compute : CreatorCompute definition end **/

/** Compute : RecieverCompute definition start **/
type RecieverCompute struct{}

func getRecieverCompute() *RecieverCompute {
	return &RecieverCompute{}
}

func (c2 *RecieverCompute) String() string {
	return "RecieverCompute"
}

func (c2 *RecieverCompute) Execute(arg *Args) {
	defer arg.WaitGroup.Done()
	packet := <-arg.Incoming
	result = packet["key"].(string)
}
/** Compute : RecieverCompute definition end **/

/** Compute : UpperCaseCompute definition start **/
type UpperCaseCompute struct{}

func getUpperCaseCompute() *UpperCaseCompute {
	return &UpperCaseCompute{}
}

func (c2 *UpperCaseCompute) String() string {
	return "UpperCaseCompute"
}

func (c2 *UpperCaseCompute) Execute(arg *Args) {
	defer arg.WaitGroup.Done()
	packet := <-arg.Incoming
	upper := strings.ToUpper(packet["key"].(string))
	packet["key"] = upper
	arg.Outgoing <- packet
}


/** tests **/
func TestLoadComputes(t *testing.T) {
	Run(getCreatorCompute(), getRecieverCompute())
	if result != "value" {
		t.Errorf("Expected %s got %s", "value", result)
	}
}

func TestUpperCaseCompute(t *testing.T) {
	Run(getCreatorCompute(), getUpperCaseCompute(), getRecieverCompute())
	if result != "VALUE" {
		t.Errorf("Expected %s got %s", "value", result)
	}
}