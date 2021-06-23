package storage

import (
	"github.com/golang/protobuf/proto"
	"gkin/message"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func NewLocalStorage() Storage {
	return &LocalStorage{}
}

/*
	本地存储
*/
type LocalStorage struct {
	localPath string
}

func (l *LocalStorage) Len(topic string) (int, error) {

	var k int
	err := filepath.Walk(l.localPath+"/"+topic, func(filename string, fi os.FileInfo, err error) error { //遍历目录
		if fi.IsDir() { // 忽略目录
			return nil
		}
		k++

		return nil

	})

	if err != nil {
		return 0, err
	}
	return k, nil
}

func (l *LocalStorage) Get(topic string, sequence int64) (*message.Message, error) {
	data, err := ioutil.ReadFile(l.localPath + "/" + topic + "/" + strconv.FormatInt(sequence, 10))
	if err != nil {
		return nil, err
	}
	msg := message.New()
	err = proto.UnmarshalMerge(data, msg)
	if err != nil {
		return nil, err
	}
	return msg, nil

}

func (l *LocalStorage) Write(msg *message.Message) error {

	msg.IsWrite = true
	data, err := proto.Marshal(msg)
	if err != nil {
		return err
	}
	err = os.MkdirAll(l.localPath+"/"+msg.GetTopic(), 0777)
	if err != nil {
		log.Println(err)
	}
	return ioutil.WriteFile(l.localPath+"/"+msg.GetTopic()+"/"+strconv.FormatInt(msg.GetSequence(), 10), data, 0777)
}

func (l *LocalStorage) IsExist(topic string, sequence int64) bool {
	panic("implement me")
}

func (l *LocalStorage) Initialization() {
	l.localPath = "./data"

}
