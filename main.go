package main

import (
	"fmt"
	oapimiddleware "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog"
	"github.com/originbenntou/modev-backend/adapter/mysql"
	"github.com/originbenntou/modev-backend/application/usecase"
	"github.com/originbenntou/modev-backend/domain/service"
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
	r := chi.NewRouter()
	r.Use(oapimiddleware.OapiRequestValidator(swagger))

	logger := httplog.NewLogger("modev-backend", httplog.Options{
		JSON: true,
	})
	r.Use(httplog.RequestLogger(logger))

	db, err := mysql.NewDB()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error db connection\n: %s", err)
		os.Exit(1)
	}

	// TODO: di container
	server := controller.NewController(
		controller.NewTweetController(
			usecase.NewTweetUseCase(
				service.NewTweetService(
					database.NewTweetRepository(db),
				),
			),
		),
	)
	gen.HandlerFromMux(server, r)

	// TODO: graceful shutdown

	// TODO: timeout setting
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		return
	}
}
