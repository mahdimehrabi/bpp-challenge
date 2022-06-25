package main

import (
	"context"
	"go.uber.org/fx"
	"net/http"
	"os"
	"time"
	"user/app/controllers"
	"user/app/infrastractures"
	"user/app/repositories"
	"user/app/routes"
	"user/app/services"
)

var BootstrapModule = fx.Options(
	infrastractures.Module,
	routes.Modules,
	controllers.Module,
	services.Module,
	repositories.Module,
	fx.Invoke(Bootstrap),
)

func Bootstrap(
	lifecycle fx.Lifecycle,
	logger infrastractures.PasargadLogger,
	userRoute *routes.UserRoute,
) {
	port := "8000"

	// create a new serve mux and register routes
	sm := http.NewServeMux()

	userRoute.Handle(sm)
	// create a new server
	s := http.Server{
		Addr:         ":" + port,        // configure the bind address
		Handler:      sm,                // set the default handler
		ErrorLog:     logger.LG,         // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	lifecycle.Append(fx.Hook{
		OnStart: func(context context.Context) error {

			// start the server
			go func() {
				logger.LG.Println("Starting server on port " + port)

				err := s.ListenAndServe()
				if err != nil {
					logger.LG.Printf("Error starting server: %s\n", err)
					os.Exit(1)
				}

			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return s.Shutdown(ctx)
		},
	})

}

func main() {
	fx.New(BootstrapModule).Run()
}
