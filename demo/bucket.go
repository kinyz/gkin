package main

import (
	"gkin/bucket"

)

func main() {

	b:=bucket.NewBucket(nil)

	b.Serve("0.0.0.0:17222")
}