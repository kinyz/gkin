package main

import (
	"gkin/kv"
	"log"
)

func main() {
	k := kv.NewMemoryKv()

	k.Set(111, "www")
	v, ok := k.Get(111)
	if ok {
		log.Println(v)
	}
	//
	//go func() {
	//	for  {
	//		time.Sleep(1*time.Second)
	//		log.Println(k.GetMap())
	//	}
	//
	//}()
	i := 0

	for {
		//time.Sleep(2*time.Second)
		i++
		log.Println(i)
		k.Set(i, i)
	}

}
