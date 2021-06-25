package main

import (
	"flag"
	"gkin"
)

func main() {
	addr := ""
	port := ""
	flag.StringVar(&addr, "addr", "0.0.0.0", "-addr addr")
	flag.StringVar(&port, "port", "17222", "-port port")
	flag.Parse()

	b := gkin.NewBucket("123", gkin.LocalStorage())

	b.Serve(addr + ":" + port)
}
