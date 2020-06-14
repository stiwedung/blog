package service

import (
	"github.com/stiwedung/blog/dao"
	"github.com/stiwedung/blog/model"
)

type articleService struct {
	lstCache   []*model.Article
	cacheIndex map[int]int
	loaded     bool
}

var articleSrv = articleService{
	cacheIndex: make(map[int]int),
}

func ArticleList() []*model.Article {
	if articleSrv.loaded {
		return articleSrv.lstCache
	}
	articleSrv.lstCache, articleSrv.loaded = dao.ArticleList()
	if articleSrv.loaded {
		for i, article := range articleSrv.lstCache {
			articleSrv.cacheIndex[article.ID] = i
		}
	}
	return articleSrv.lstCache
}

func ShowArticle(id int) *model.Article {
	if !articleSrv.loaded {
		ArticleList()
	}
	if !articleSrv.loaded {
		return nil
	}
	idx, ok := articleSrv.cacheIndex[id]
	if !ok {
		return nil
	}
	return articleSrv.lstCache[idx]
}
