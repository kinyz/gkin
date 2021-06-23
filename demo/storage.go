package main

import (
	"gkin/message"
	"gkin/storage"
	"log"
	"time"
)

func main() {
	s := storage.NewLocalStorage()

	s.Initialization()
	msg := message.New()

	msg.SetTopic("myTopic2")
	msg.SetTimesTamp(time.Now().UnixNano())
	msg.SetKey("im key")
	msg.SetValue([]byte("im value"))

	for i := 0; i < 5; i++ {
		msg.Sequence = int64(i)
		err := s.Write(msg)
		if err != nil {
			log.Println(err)
		}
	}

	message.Pool.Put(msg)

	for i := 0; i < 5; i++ {
		get, err := s.Get("myTopic2", int64(i))
		if err != nil {
			log.Println(err)
			continue
		}
		message.Pool.Put(get)
		log.Println(get)
		//get.Put()
	}

	i, err := s.Len("myTopic2")
	if err != nil {
		panic(err)
	}
	log.Println(i)

}
