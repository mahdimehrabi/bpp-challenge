package controllers

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"user/app/infrastractures"
	"user/app/interfaces"
	"user/app/models"
	"user/app/services"
)

type UserController struct {
	logger      interfaces.Logger
	userService *services.UserService
}

func NewUserController(logger infrastractures.PasargadLogger, userService *services.UserService) UserController {
	return UserController{
		logger:      &logger,
		userService: userService,
	}
}

func (c UserController) Create() endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(models.UserRequest)
		err := c.userService.Create(req.Name, req.VIP)
		if err != nil {
			return models.OperationMessage{"failed to create user!"}, nil
		}

		return models.OperationMessage{"User created successfully!"}, nil
	}
}

func (c UserController) List() endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		users, err := c.userService.List()
		if err != nil {
			return users, nil
		}
		return users, nil
	}
}

func (c UserController) Detail() endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(models.FindUserRequest)
		user, err := c.userService.Detail(req.ID)
		if err != nil {
			return user, nil
		}
		return user, nil
	}
}

func (c UserController) Update() endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(models.UpdateUserRequest)
		err := c.userService.Update(req.ID, req.Name, req.VIP)
		if err != nil {
			return models.OperationMessage{"failed to update user!"}, nil
		}

		return models.OperationMessage{"User updated successfully!"}, nil
	}
}

func (c UserController) Delete() endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(models.FindUserRequest)
		err := c.userService.Delete(req.ID)
		if err != nil {
			return models.OperationMessage{"failed to delete user!"}, nil
		}

		return models.OperationMessage{"User delete successfully!"}, nil
	}
}
