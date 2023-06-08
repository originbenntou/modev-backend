package main

import (
	"fmt"
	oapiMiddleware "github.com/deepmap/oapi-codegen/pkg/middleware"
	oapi "github.com/originbenntou/modev-backend/generated"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type apiController struct{}

func (a apiController) GetTweets(ctx echo.Context, params oapi.GetTweetsParams) error {
	return ctx.JSON(http.StatusOK, &oapi.Tweets{
		Id:      1,
		AddDate: "2022-01-01",
		Url:     "https://aaaa.com",
		Tags: &[]string{
			"Golang",
			"Java",
		},
	})
}

func main() {
	e := echo.New()

	swagger, err := oapi.GetSwagger()
	if err != nil {
		panic(err)
	}
	swagger.Servers = nil
	e.Use(oapiMiddleware.OapiRequestValidator(swagger))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	api := apiController{}
	oapi.RegisterHandlers(e, api)

	for _, route := range e.Routes() {
		fmt.Printf("%s %s\n", route.Method, route.Path)
	}

	e.Logger.Fatal(e.Start(":8080"))
}
