package storage

type Storage interface {
	Initialization()
	Get(path ,name string)(interface{},error)
	Write(path ,name string,data []byte)error
	IsExist(path ,name string)bool
}
