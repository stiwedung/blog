package dao

import (
	"github.com/stiwedung/blog/model"
	"github.com/stiwedung/libgo/log"
)

func UserLogin(userName, passwd string) (*model.User, bool) {
	ret := &model.User{}
	if _, err := db.Where("user_name=?", userName).Select("`id`, `user_name`, `passcode`, `passwd`, `is_admin`").Get(ret); err != nil {
		log.Errorf("User not create: %v", err)
		return ret, false
	}
	return ret, true
}

func CreateUser(userName, passcode, passwd, loginIP string, isAdmin bool) bool {
	user := &model.User{
		UserName: userName,
		Passcode: passcode,
		Passwd:   passwd,
		LoginIP:  loginIP,
		IsAdmin:  isAdmin,
	}
	if _, err := db.InsertOne(user); err != nil {
		log.Errorf("create user failed: %v", err)
		return false
	}
	return true
}
