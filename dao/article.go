package dao

import (
	"github.com/stiwedung/blog/model"
	"github.com/stiwedung/libgo/log"
)

func ArticleList() ([]*model.Article, bool) {
	var ret []*model.Article
	if err := db.Select(`id, title, content, create_time`).Find(&ret); err != nil {
		log.Errorf("load all article info form database error: %v", err)
		return nil, false
	}
	return ret, true
}
