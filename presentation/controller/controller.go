package controller

import (
	"encoding/json"
	"net/http"
)

type Controller struct {
	*TweetController
}

func NewController(t *TweetController) *Controller {
	return &Controller{
		TweetController: t,
	}
}

func RespondOK(w http.ResponseWriter, result any) {
	err := json.NewEncoder(w).Encode(result)

	if err != nil {
		// RespondError(w, errors.Wrap(err, "faile to respond_ok"))
		return
	}
}
