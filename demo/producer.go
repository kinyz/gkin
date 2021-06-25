package main

import (
	"gkin"
	"gkin/bucket"
	"gkin/utils"
	"log"
	"strconv"
	"time"
)

func main() {

	p, err := gkin.NewProducer("127.0.0.1:17222", utils.NewUuid(), bucket.StreamKey)
	//p, err := producer.NewProducer("127.0.0.1:17222", utils.NewUuid(), bucket.StreamKey)
	if err != nil {
		panic(err)
	}
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		err := p.ASyncSend("user", "im key", []byte("我是value"+strconv.Itoa(i)))
		if err != nil {
			log.Println(err)
		}
		log.Println("发送第", i)
	}
	select {}

}
