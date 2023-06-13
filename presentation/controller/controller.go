package controller

import (
	"encoding/json"
	"github.com/go-chi/httplog"
	"github.com/originbenntou/modev-backend/presentation/failure"
	"github.com/pkg/errors"
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

func RespondOK(w http.ResponseWriter, r *http.Request, result any) {
	_ = httplog.LogEntry(r.Context())

	err := json.NewEncoder(w).Encode(result)
	if err != nil {
		RespondError(w, r, errors.Wrap(err, "faile to respond_ok"))
	}
	return
}

func RespondError(w http.ResponseWriter, r *http.Request, err error) {
	var appError *failure.AppError
	if errors.As(err, &appError) {
		// TODO: log info
		switch appError.Code {
		//case failure.ErrInvalid:
		//	w.WriteHeader(http.StatusBadRequest)
		//case failure.ErrForbidden:
		//	w.WriteHeader(http.StatusForbidden)
		//case failure.ErrNotFound:
		//	w.WriteHeader(http.StatusNotFound)
		//case failure.ErrConflict:
		//	w.WriteHeader(http.StatusConflict)
		default:
			// TODO: log unknown code
			w.WriteHeader(http.StatusInternalServerError)
			oplog := httplog.LogEntry(r.Context())
			oplog.Error().Msg(err.Error())
		}
		return
	}

	// TODO: log fatal
	w.WriteHeader(http.StatusInternalServerError)
	oplog := httplog.LogEntry(r.Context())
	oplog.Error().Msg(err.Error())
	return
}
