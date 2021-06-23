package main

import (
	"gkin/bucket"
	"gkin/storage"
)

func main() {

	b := bucket.NewBucket(storage.NewLocalStorage())

	b.Serve("0.0.0.0:17222")
}
