package main

import (
	"fmt"
	middleware "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
	"github.com/go-chi/chi/v5"
	"github.com/originbenntou/modev-backend/application/service"
	"github.com/originbenntou/modev-backend/application/usecase"
	"github.com/originbenntou/modev-backend/gen"
	"github.com/originbenntou/modev-backend/infrastructure/database"
	"github.com/originbenntou/modev-backend/presentation/controller"
	"net/http"
	"os"
)

func main() {
	swagger, err := gen.GetSwagger()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	swagger.Servers = nil
	router := chi.NewRouter()
	router.Use(middleware.OapiRequestValidator(swagger))

	// TODO: di container
	server := controller.NewController(
		controller.NewTweetController(
			usecase.NewTweetUseCase(
				service.NewTweetService(
					database.NewTweetRepository(),
				),
			),
		),
	)
	gen.HandlerFromMux(server, router)

	// TODO: graceful shutdown

	// TODO: timeout setting
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		return
	}
}
