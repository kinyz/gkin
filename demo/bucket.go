package main

import (
	"flag"
	"gkin/bucket"
	"gkin/storage"
)

func main() {
	addr := ""
	port := ""
	flag.StringVar(&addr, "addr", "0.0.0.0", "-addr addr")
	flag.StringVar(&port, "port", "17222", "-port port")
	flag.Parse()

	b := bucket.NewBucket(storage.NewLocalStorage())

	b.Serve(addr + ":" + port)
}
