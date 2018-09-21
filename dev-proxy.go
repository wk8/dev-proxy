package main

import (
	"flag"
	"fmt"
)

var VERSION string

func main() {
	// parse command line args
	listenPort := flag.Int("l", 8282, "The local port to listen on")
	target := flag.String("t", "http://google.com", "The target to proxy to")
	isVersion := flag.Bool("v", false, "Outputs the version then exits")
	flag.Parse()

	if *isVersion {
		fmt.Println("Version:", VERSION)
		return
	}

	// start the proxy
	proxy := NewProxy(*target, nil)
	proxy.Start(*listenPort, true)
}
