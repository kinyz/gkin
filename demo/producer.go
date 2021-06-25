package main

import (
	"gkin"
	"gkin/utils"
	"log"
	"strconv"
)

func main() {

	p, err := gkin.NewProducer("127.0.0.1:17222", utils.NewUuid(), "")
	//p, err := producer.NewProducer("127.0.0.1:17222", utils.NewUuid(), bucket.StreamKey)
	if err != nil {
		panic(err)
	}

	for i := 1; i < 101; i++ {
		//time.Sleep(1 * time.Second)
		b, err := p.SyncSend("user", "im key", []byte("我是value"+strconv.Itoa(i)))
		if err != nil {
			log.Println(err)
		}
		log.Println("发送第", i, b)
	}
	select {}

}
