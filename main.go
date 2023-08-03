package main

import (
	"fmt"
	oapimiddleware "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
	"github.com/go-chi/chi/v5"
	"github.com/originbenntou/modev-backend/adapter/mysql"
	"github.com/originbenntou/modev-backend/application/usecase"
	"github.com/originbenntou/modev-backend/common/logger"
	"github.com/originbenntou/modev-backend/domain/service"
	"github.com/originbenntou/modev-backend/gen"
	"github.com/originbenntou/modev-backend/infrastructure/database"
	"github.com/originbenntou/modev-backend/presentation/controller"
	"github.com/originbenntou/modev-backend/presentation/middleware"
	"golang.org/x/exp/slog"
	"log"
	"net/http"
	"os"
)

func main() {
	// TODO: app config を切り出す

	swagger, err := gen.GetSwagger()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	swagger.Servers = nil
	r := chi.NewRouter()
	r.Use(oapimiddleware.OapiRequestValidator(swagger))

	l := logger.New(logger.Opts{
		Level: slog.LevelDebug,
		OnError: func(l *logger.Logger, msg string, err error, arg ...any) {
			traceIDContext, ok := l.LoggerContext("traceID")
			if !ok {
				log.Println(msg)
				return
			}

			log.Printf("%s のリクエストでエラーが起きたよ\n", traceIDContext.Value)
		},
	})
	r.Use(middleware.LoggerInjector(l, "modev-backend"))

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
