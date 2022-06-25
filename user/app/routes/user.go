package routes

import (
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
	sm.Handle("/users", httptransport.NewServer(
		r.userController.Create(),
		models.DecodeUserRequest,
		models.EncodeResponse,
	))
}

func NewUserRoute(logger infrastractures.PasargadLogger, userController controllers.UserController) *UserRoute {
	return &UserRoute{logger: logger, userController: userController}
}
