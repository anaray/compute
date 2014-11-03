package compute


type Args struct {
	Incoming chan string
	Outgoing chan string
}

type Plugins interface {
	Execute(Args)
}


func Run(plugins ...Plugins) {
	//done := make()
	in := make(chan string, 100)
	var indx = 1
	for _, plugin := range plugins {
		out := make(chan string, 100)
		arg := Args{Incoming: in, Outgoing: out}
		for i := 0; i < indx; i++ {
			go plugin.Execute(arg)
		}
		in = out
		indx += 1
	}


	for { //i := 0; i < 100; i++ {
		_=<-in
		//fmt.Println(output)
	}

}