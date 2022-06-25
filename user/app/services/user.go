package services

import (
	"user/app/interfaces"
	"user/app/repositories"
)

type UserService struct {
	logger         interfaces.Logger
	userRepository repositories.UserRepository
}

func (s UserService) Create(name string, vip bool) error {
	err := s.userRepository.Create(name, vip)
	if err != nil {
		return nil
	}
	s.logger.Error(err.Error())
	return err
}
