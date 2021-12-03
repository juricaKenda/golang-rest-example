package repository

import (
	"github.com/juricaKenda/golang-rest-example/src/model"
	"github.com/juricaKenda/golang-rest-example/src/server"
)

type Repo struct {
	store []model.Article
	log   server.ILogger
}

func NewArticleRepo(log server.ILogger) server.IArticleRepository {
	i := &Repo{
		store: make([]model.Article, 0),
		log:   log,
	}
	i.init()
	return i
}

func (i *Repo) GetAllArticles() ([]model.Article, error) {
	i.log.Trace("getting all articles from the repo, im inside the repo")
	return i.store, nil
}

func (i *Repo) init() {
	i.log.Trace("init repo..")
	i.store = append(i.store, model.Article{
		Author: "Rakesh",
		Title:  "Why do we need all these interfaces??",
		Body:   "Please help me understand. Let's build faster!!",
	})

	i.store = append(i.store, model.Article{
		Author: "Jurica",
		Title:  "We are building configurable software",
		Body:   "Quality > speed",
	})

}
