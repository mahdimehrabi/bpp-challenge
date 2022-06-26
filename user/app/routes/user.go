package routes

import (
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"net/http"
	"user/app/controllers"
	"user/app/infrastractures"
	"user/app/models"
)

type UserRoute struct {
	logger         infrastractures.PasargadLogger
	userController controllers.UserController
}

func (r UserRoute) Handle(sm *http.ServeMux) {
	sm.Handle("/users", Handle(r.userController.Create()))
	sm.Handle("/users/list", Handle(r.userController.List()))
}

func NewUserRoute(logger infrastractures.PasargadLogger, userController controllers.UserController) *UserRoute {
	return &UserRoute{logger: logger, userController: userController}
}

func Handle(endpoint endpoint.Endpoint) *httptransport.Server {
	return httptransport.NewServer(
		endpoint,
		models.DecodeUserRequest,
		models.EncodeResponse,
	)
}
