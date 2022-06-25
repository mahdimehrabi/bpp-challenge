package routes

import (
	"fmt"
	"net/http"
	"user/app/infrastractures"
)

type UserRoute struct {
	logger infrastractures.PasargadLogger
}

func (r UserRoute) Handle(sm *http.ServeMux) {
	sm.HandleFunc("/users", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("ooooooooooook")
	})
}

func NewUserRoute(logger infrastractures.PasargadLogger) *UserRoute {
	return &UserRoute{logger: logger}
}
