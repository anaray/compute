compute
=======
(status: initial commits, work-in-progress)
Create computes, chain them, in order to a create complex compuation workflow

A simple example: chaining udplistener -> uppercase -> stdout 
```
package main

import ("compute")

func main() {
	compute.Run(UDPListener(), UpperCase(), Stdout())
}
```

another example: chaining twitter_listener -> sentiment_analyzer -> push_to_visualizer 
```
package main

import ("compute")

func main() {
	compute.Run(TwitterListener(), SentimentAnalyzer(), Visualize())
}
```

Design Goals:

1. capability to create/initialize computes in different physical machines.
2. capability to create multiple instances of same computes in order to accomodate load(fan-in or fan-out).
3. Need a Consensus algorithm support (RAFT https://raftconsensus.github.io/) when working in distributed mode with multiple instances ? 

Writing a compute:

1. Define a Compute struct.
2. Have a initializer function.
3. Have the struct implement Computes interface.
4. Args argument in Execute method provides.
	1. Incoming channel.
	2. Outgoing channel.
	2. a WaitGroup.
	3. Container to put configurations.
	4. a Logger.

For example: https://github.com/anaray/compute/blob/master/compute_test.go
