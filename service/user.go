package service

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"time"

	"github.com/stiwedung/blog/dao"
	"github.com/stiwedung/blog/model"
)

type userService struct {
	users map[string]*model.User
	rand  *rand.Rand
}

var userSrv = userService{
	users: make(map[string]*model.User),
	rand:  rand.New(rand.NewSource(time.Hour.Nanoseconds())),
}

var (
	ErrLogin      = errors.New("User Login failed")
	ErrPasswd     = errors.New("User password error")
	ErrCreateUser = errors.New("User create failed")
)

func checkUser(passwd string, user *model.User) error {
	hashMd5 := md5.New()
	io.WriteString(hashMd5, user.Passcode+passwd)
	checkPasswd := fmt.Sprintf("%x", hashMd5.Sum(nil))
	if user.Passwd != checkPasswd {
		return ErrPasswd
	}
	return nil
}

func Login(username, passwd string, isLocal bool) error {
	if user, ok := userSrv.users[username]; ok {
		return checkUser(passwd, user)
	}
	if user, ok := dao.GetUser(username); ok {
		userSrv.users[username] = user
		return checkUser(passwd, user)
	}
	if !isLocal {
		return ErrLogin
	}
	passcode := fmt.Sprintf("%x", userSrv.rand.Int31())
	hashMd5 := md5.New()
	io.WriteString(hashMd5, passcode+passwd)
	newPasswd := fmt.Sprintf("%x", hashMd5.Sum(nil))
	user, ok := dao.CreateUser(username, passcode, newPasswd, true)
	if !ok {
		return ErrCreateUser
	}
	userSrv.users[username] = user
	return nil
}
