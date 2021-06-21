package message

import (
	"encoding/json"
	"google.golang.org/protobuf/proto"
)

func (x *Message) Put() {
	Pool.Put(x)
}

func (x *Message) SetKey(key string) {
	x.Key = key

}
func (x *Message) SetValue(value []byte) {
	x.Value = value
}
func (x *Message) SetTopic(topic string) {
	x.Topic = topic
}
func (x *Message) SetHeaders(headers map[string]string) {
	x.Headers = headers
}
func (x *Message) SetTimesTamp(time int64) {
	x.TimesTamp = time
}

func (x *Message) GetHeader(key string) string {
	h, ok := x.Headers[key]
	if !ok {
		return ""
	}
	return h
}
func (x *Message) SetHeader(key, value string) {
	x.Headers[key] = value
}

func (x *Message) ReadProtobuf(ptr proto.Message) error {
	return proto.Unmarshal(x.GetValue(), ptr)
}
func (x *Message) ReadJson(ptr interface{}) error {
	return json.Unmarshal(x.GetValue(), ptr)
}
func (x *Message) ReadString() string {
	return string(x.GetValue())
}
func (x *Message) ReadBytes() []byte {
	return x.GetValue()
}

func (x *Message) SetProducer(name string) {
	x.Producer = name

}
