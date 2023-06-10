package main

import (
	"encoding/json"
	"fmt"
	middleware "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
	"github.com/go-chi/chi/v5"
	"github.com/originbenntou/modev-backend/gen"
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
	gen.HandlerFromMux(NewServer(), router)

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		return
	}
}

type Server struct{}

func (s *Server) GetTweets(w http.ResponseWriter, r *http.Request, p gen.GetTweetsParams) {
	w.WriteHeader(http.StatusOK)
	var err = json.NewEncoder(w).Encode([]gen.Tweet{
		{
			AddDate: "2000-01-01",
			Id:      1,
			Tags:    nil,
			Url:     "",
		},
		{
			AddDate: "2000-01-02",
			Id:      2,
			Tags:    nil,
			Url:     "",
		},
	})
	if err != nil {
		return
	}
}

func NewServer() *Server {
	return &Server{}
}
