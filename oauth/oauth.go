package oauth

import (
	"gkin/pb"
)

type Oauth interface {
	Connection(connect pb.Connect) error
}
