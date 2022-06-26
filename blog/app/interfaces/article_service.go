package interfaces

import (
	"blog/app/models"
)

type ArticleService interface {
	Create(string, bool) error
	List() ([]models.Article, error)
	Detail() (models.Article, error)
	Delete(int64) error
	Update(string, bool) error
}
