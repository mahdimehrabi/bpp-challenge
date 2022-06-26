package repositories

import (
	"context"
	infrastractures2 "user/app/infrastractures"
	interfaces2 "user/app/interfaces"
	"user/app/models"
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

func (u *UserRepository) List() (users []models.User, err error) {
	values, err := u.db.Query(context.Background(),
		"SELECT * FROM users",
		[]interface{}{})
	for _, v := range values {
		user := models.User{
			ID:   int64(v[0].(int32)),
			VIP:  v[1].(bool),
			Name: v[2].(string),
		}
		users = append(users, user)
	}
	return users, err
}

func (u *UserRepository) Detail(id int64) (user models.User, err error) {
	err = u.db.QueryRow(context.Background(),
		"SELECT id,name,vip FROM users WHERE id=$1",
		[]interface{}{id}, &user.ID, &user.Name, &user.VIP)
	return
}

func (u *UserRepository) Update(id int64, name string, vip bool) error {
	_, err := u.db.Exec(context.Background(),
		"UPDATE users SET name=$1,vip=$2 WHERE id=$3",
		[]interface{}{name, vip, id})
	return err
}
