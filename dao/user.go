package dao

import (
	"github.com/stiwedung/blog/model"
	"github.com/stiwedung/libgo/log"
)

func GetUser(userName string) (*model.User, bool) {
	ret := &model.User{}
	if ok, err := db.Where("user_name=?", userName).Get(ret); err != nil {
		log.Errorf("User not create: %v", err)
		return ret, false
	} else if !ok {
		return ret, false
	}
	return ret, true
}

func CreateUser(userName, passcode, passwd string, isAdmin bool) (*model.User, bool) {
	user := &model.User{
		UserName: userName,
		Passcode: passcode,
		Passwd:   passwd,
		IsAdmin:  isAdmin,
	}
	if _, err := db.InsertOne(user); err != nil {
		log.Errorf("create user failed: %v", err)
		return user, false
	}
	return user, true
}
