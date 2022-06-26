package routes

import (
	"blog/app/controllers"
	"blog/app/infrastractures"
	"blog/app/middlewares"
	"blog/app/models"
	"context"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"net/http"
)

type request func(context.Context, *http.Request) (interface{}, error)

type ArticleRoute struct {
	logger         infrastractures.PasargadLogger
	blogController controllers.BlogController
	vipMiddleware  middlewares.VipMiddleware
}

func (r ArticleRoute) Handle(sm *http.ServeMux) {
	sm.Handle("/articles", Handle(r.blogController.Create(), models.DecodeArticleRequest))
	sm.Handle("/articles/list", Handle(r.blogController.List(), models.DecodeArticleRequest))

	sm.Handle("/articles/detail", Handle(r.blogController.Detail(),
		r.vipMiddleware.Handle(models.DecodeFindArticleRequest),
	))
	sm.Handle("/articles/update", Handle(r.blogController.Update(),
		r.vipMiddleware.Handle(models.DecodeUpdateArticleRequest),
	))
	sm.Handle("/articles/delete", Handle(r.blogController.Delete(),
		r.vipMiddleware.Handle(models.DecodeFindArticleRequest),
	))
}

func NewUserRoute(logger infrastractures.PasargadLogger,
	userController controllers.BlogController,
	vipMiddleware middlewares.VipMiddleware) *ArticleRoute {
	return &ArticleRoute{
		logger:         logger,
		blogController: userController,
		vipMiddleware:  vipMiddleware,
	}
}

func Handle(endpoint endpoint.Endpoint, requestFunc httptransport.DecodeRequestFunc) *httptransport.Server {
	return httptransport.NewServer(
		endpoint,
		requestFunc,
		models.EncodeResponse,
	)
}
