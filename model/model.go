package model

import "time"

type (
	Article struct {
		ID         int       `xorm:"pk autoincr"`
		Title      string    `xorm:"notnull"`
		Content    string    `xorm:"notnull"`
		CreateTime time.Time `xorm:"notnull created"`
		UpdateTime time.Time `xorm:"notnull updated"`
		Delete     bool      `xorm:"notnull default(0)"`
	}

	Tag struct {
		ID        int    `xorm:"pk autoincr"`
		Tag       string `xorm:"index notnull"`
		ArticleID int    `xorm:"index notnull"`
	}

	User struct {
		ID       int    `xorm:"pk autoincr"`
		UserName string `xorm:"index notnull"`
		Passcode string `xorm:"notnull"`
		Passwd   string `xorm:"notnull"`
		LoginIP  string `xorm:"notnull"`
		IsAdmin  bool   `xorm:"notnull default(0)"`
	}
)

var Models = []interface{}{
	Article{}, Tag{}, User{},
}
