package main

import (
	"flag"
	"net"
	"os"
	"fmt"
)


var verbose = flag.Bool("v", false, "verbose")
var service = flag.String("s", "", "service")
var proto = flag.String("p", "", "proto")

func main () {
	flag.Parse()

	args := flag.Args()

	if len(args) != 1 {
		fmt.Println("Please specify a single record to lookup")
		os.Exit(1)
	}

	if *verbose {
		fmt.Printf("looking up %s\n", args[0])
	}

	cname, addrs, err := net.LookupSRV(*service, *proto, args[0])

	if err != nil {
		fmt.Printf("unable to lookup %s: %s\n", args[0], err)
		os.Exit(2)
	}

	if *verbose {
		fmt.Printf("cname: %s\n", cname)
	}

	for _, addr := range(addrs) {
		fmt.Printf("%s:%d\n", addr.Target, addr.Port)
	}

}
