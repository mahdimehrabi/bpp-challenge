package repositories

import (
	"context"
	infrastractures2 "user/app/infrastractures"
	interfaces2 "user/app/interfaces"
)

type UserRepository struct {
	logger interfaces2.Logger
	db     interfaces2.DB
}

func NewUserRepository(logger infrastractures2.PasargadLogger, db infrastractures2.PgxDB) UserRepository {
	return UserRepository{
		logger: &logger,
		db:     &db,
	}
}

func (u *UserRepository) Create(name string, vip bool) error {
	_, err := u.db.Exec(context.Background(),
		"INSERT INTO users(name,vip) VALUES($1,$2)",
		[]interface{}{name, vip})
	return err
}
