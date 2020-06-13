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
)

var Models = []interface{}{
	Article{}, Tag{},
}
