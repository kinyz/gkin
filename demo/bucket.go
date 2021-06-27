package main

import (
	"flag"
	"github.com/kinyz/gkin"
)

func main() {
	addr := ""
	port := ""
	key := ""
	flag.StringVar(&addr, "addr", "0.0.0.0", "-addr addr")
	flag.StringVar(&port, "port", "17222", "-port port")
	flag.StringVar(&key, "key", "", "-key key")
	flag.Parse()

	b := gkin.NewBucket(key, gkin.LocalStorage())

	b.Serve(addr + ":" + port)
}
