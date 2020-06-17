package controller

import "encoding/gob"

const (
	userInfo = "user_info"
)

func init() {
	gob.Register(sessionData{})
}

type sessionData struct {
	UserName string
}
