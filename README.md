compute
=======
(status: initial commits, work-in-progress)
Create computes, chain them, in order to a create complex compuation workflow

A simple example: chaining udplistener -> uppercase -> stdout 
```
package main

import ("plugins")

func main() {
	plugins.Run(plugins.UDPListener(), plugins.UpperCase(), plugins.Stdout())
}
```

Design Goals:
1) capability to create/initialize computes in different physical machines.
2) capability to create multiple instances of same computes in order to accomodate load(fan-in or fan-out).
