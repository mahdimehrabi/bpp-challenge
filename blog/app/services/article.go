package services

import (
	"user/app/infrastractures"
	"user/app/interfaces"
	"user/app/models"
	"user/app/repositories"
)

type ArticleService struct {
	logger            interfaces.Logger
	articleRepository repositories.ArticleRepository
}

func (s ArticleService) Create(title string, body string) error {
	err := s.articleRepository.Create(title, body)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	return nil
}

func (s ArticleService) List() ([]models.Article, error) {
	articles, err := s.articleRepository.List()
	if err != nil {
		s.logger.Error(err.Error())
		return articles, err
	}
	return articles, nil
}

func (s ArticleService) Detail(id int64) (article models.Article, err error) {
	article, err = s.articleRepository.Detail(id)
	if err != nil {
		s.logger.Error(err.Error())
		return
	}
	return
}

func (s ArticleService) Update(id int64, title string, body string) error {
	err := s.articleRepository.Update(id, title, body)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	return nil
}

func (s ArticleService) Delete(id int64) error {
	err := s.articleRepository.Delete(id)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	return nil
}

func NewUserService(logger infrastractures.PasargadLogger, userRepository repositories.ArticleRepository) *ArticleService {
	return &ArticleService{
		logger:            &logger,
		articleRepository: userRepository,
	}
}
