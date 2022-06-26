package routes

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"net/http"
	"user/app/controllers"
	"user/app/infrastractures"
	"user/app/models"
)

type request func(context.Context, *http.Request) (interface{}, error)

type UserRoute struct {
	logger         infrastractures.PasargadLogger
	userController controllers.UserController
}

func (r UserRoute) Handle(sm *http.ServeMux) {
	sm.Handle("/users", Handle(r.userController.Create(), models.DecodeUserRequest))
	sm.Handle("/users/list", Handle(r.userController.List(), models.DecodeUserRequest))
	sm.Handle("/users/detail", Handle(r.userController.Detail(), models.DecodeFindUserRequest))
}

func NewUserRoute(logger infrastractures.PasargadLogger, userController controllers.UserController) *UserRoute {
	return &UserRoute{logger: logger, userController: userController}
}

func Handle(endpoint endpoint.Endpoint, requestFunc httptransport.DecodeRequestFunc) *httptransport.Server {
	return httptransport.NewServer(
		endpoint,
		requestFunc,
		models.EncodeResponse,
	)
}
