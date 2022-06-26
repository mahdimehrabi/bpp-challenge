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

type ArticleRoute struct {
	logger         infrastractures.PasargadLogger
	blogController controllers.BlogController
}

func (r ArticleRoute) Handle(sm *http.ServeMux) {
	sm.Handle("/articles", Handle(r.blogController.Create(), models.DecodeArticleRequest))
	sm.Handle("/articles/list", Handle(r.blogController.List(), models.DecodeArticleRequest))
	sm.Handle("/articles/detail", Handle(r.blogController.Detail(), models.DecodeFindArticleRequest))
	sm.Handle("/articles/update", Handle(r.blogController.Update(), models.DecodeUpdateArticleRequest))
	sm.Handle("/articles/delete", Handle(r.blogController.Delete(), models.DecodeFindArticleRequest))
}

func NewUserRoute(logger infrastractures.PasargadLogger, userController controllers.BlogController) *ArticleRoute {
	return &ArticleRoute{logger: logger, blogController: userController}
}

func Handle(endpoint endpoint.Endpoint, requestFunc httptransport.DecodeRequestFunc) *httptransport.Server {
	return httptransport.NewServer(
		endpoint,
		requestFunc,
		models.EncodeResponse,
	)
}
