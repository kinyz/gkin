package oauth

import (
	"gkin/connect"
)

type Oauth interface {
	Connection(connect connect.Connect) error
}
