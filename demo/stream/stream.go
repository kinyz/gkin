package main

import "gkin/stream"

func main() {

	s := stream.NewStream()
	s.Serve("0.0.0.0:19999")
}
