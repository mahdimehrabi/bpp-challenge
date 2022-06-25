package interfaces

import (
	"user/app/models"
)

type UserService interface {
	Create(string, bool) error
	List() ([]models.User, error)
	Detail() (models.User, error)
	Delete(int64) error
	Update(string, bool) error
}
