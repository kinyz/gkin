package main

import (
	"gkin"
	"log"
)

func main() {

	cli := gkin.NewConsumer("1.116.248.126:17222", "123")

	go cli.Watch("user", "eee", test)
	go cli.Watch("user", "eee", test2)

	select {}
}

func test(message gkin.Message) {

	log.Println("test1", message)
}

func test2(message gkin.Message) {

	log.Println("test2", message)
}
