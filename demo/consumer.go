package main

import (
	"github.com/kinyz/gkin"
	"log"
	"time"
)

func main() {

	cli := gkin.NewConsumer("127.0.0.1:17222", "123")

	t := &Test{}
	go cli.Watch("user", "eee", t.test)
	go cli.Watches([]string{"user", "ddddd"}, "eee1", test2)

	select {}
}

type Test struct {
	start int64
	b     bool
}

func (t *Test) test(message gkin.Message) {
	if !t.b {
		t.b = true
		t.start = time.Now().UnixNano()
	}
	log.Println("test1", message)
	if message.GetSequence() == 100000 {
		log.Println("共耗时:", time.Now().UnixNano()-t.start)
	}
}

func test2(message gkin.Message) {

	log.Println("test2", message)
}
