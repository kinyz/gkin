package utils

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/satori/go.uuid"
)

// NewToken 新建随机token
func NewToken() string {
	h := md5.New()
	h.Write([]byte(NewUuid()))
	return hex.EncodeToString(h.Sum(nil))
}

// NewUuid 新建随机uuid4
func NewUuid() string {
	// uuid.Must(uuid.NewV4())
	return uuid.NewV4().String()
}

