package controllers

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"user/app/infrastractures"
	"user/app/interfaces"
	"user/app/models"
	"user/app/services"
)

type BlogController struct {
	logger         interfaces.Logger
	articleService *services.ArticleService
}

func NewBlogController(logger infrastractures.PasargadLogger, userService *services.ArticleService) BlogController {
	return BlogController{
		logger:         &logger,
		articleService: userService,
	}
}

func (c BlogController) Create() endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(models.ArticleRequest)
		err := c.articleService.Create(req.Title, req.Body)
		if err != nil {
			return models.OperationMessage{"failed to create usearticler!"}, nil
		}

		return models.OperationMessage{"Article created successfully!"}, nil
	}
}

func (c BlogController) List() endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		users, err := c.articleService.List()
		if err != nil {
			return users, nil
		}
		return users, nil
	}
}

func (c BlogController) Detail() endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(models.FindArticleResponse)
		user, err := c.articleService.Detail(req.ID)
		if err != nil {
			return models.OperationMessage{Message: "No article found"}, nil
		}
		return user, nil
	}
}

func (c BlogController) Update() endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(models.UpdateArticleRequest)
		err := c.articleService.Update(req.ID, req.Title, req.Body)
		if err != nil {
			return models.OperationMessage{"failed to update usearticler!"}, nil
		}

		return models.OperationMessage{"Article updated successfully!"}, nil
	}
}

func (c BlogController) Delete() endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(models.FindArticleResponse)
		err := c.articleService.Delete(req.ID)
		if err != nil {
			return models.OperationMessage{"failed to delete article!"}, nil
		}

		return models.OperationMessage{"Article delete successfully!"}, nil
	}
}
