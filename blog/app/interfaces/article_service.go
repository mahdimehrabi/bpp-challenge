package interfaces

import (
	"user/app/models"
)

type UserService interface {
	Create(string, bool) error
	List() ([]models.Article, error)
	Detail() (models.Article, error)
	Delete(int64) error
	Update(string, bool) error
}
