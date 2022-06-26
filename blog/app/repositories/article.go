package repositories

import (
	infrastractures2 "blog/app/infrastractures"
	interfaces2 "blog/app/interfaces"
	"blog/app/models"
	"context"
)

type ArticleRepository struct {
	logger interfaces2.Logger
	db     interfaces2.DB
}

func NewUserRepository(logger infrastractures2.PasargadLogger, db infrastractures2.PgxDB) ArticleRepository {
	return ArticleRepository{
		logger: &logger,
		db:     &db,
	}
}

func (u *ArticleRepository) Create(title string, body string) error {
	_, err := u.db.Exec(context.Background(),
		"INSERT INTO articles(title,body) VALUES($1,$2)",
		[]interface{}{title, body})
	return err
}

func (u *ArticleRepository) List() (articles []models.Article, err error) {
	values, err := u.db.Query(context.Background(),
		"SELECT * FROM articles",
		[]interface{}{})
	for _, v := range values {
		article := models.Article{
			ID:    int64(v[0].(int32)),
			Title: v[1].(string),
			Body:  v[2].(string),
		}
		articles = append(articles, article)
	}
	return articles, err
}

func (u *ArticleRepository) Detail(id int64) (article models.Article, err error) {
	err = u.db.QueryRow(context.Background(),
		"SELECT id,title,body FROM articles WHERE id=$1",
		[]interface{}{id}, &article.ID, &article.Title, &article.Body)
	return
}

func (u *ArticleRepository) Update(id int64, title string, body string) error {
	_, err := u.db.Exec(context.Background(),
		"UPDATE articles SET title=$1,body=$2 WHERE id=$3",
		[]interface{}{title, body, id})
	return err
}

func (u *ArticleRepository) Delete(id int64) error {
	_, err := u.db.Exec(context.Background(),
		"DELETE FROM articles WHERE id=$1",
		[]interface{}{id})
	return err
}
