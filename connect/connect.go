package connect

type Connect interface {
	GetClientId() string
	GetKey() string
	GetToken() string
	IsOauth() bool
	Close()
}

type connect struct {
	Connection
}

func (x *Connection) IsOauth() bool {
	return x.GetOauth()
}

func (x *Connection) Close() {

}
