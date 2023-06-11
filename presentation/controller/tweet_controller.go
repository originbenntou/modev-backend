package controller

import (
	"context"
	"github.com/originbenntou/modev-backend/application/usecase"
	"github.com/originbenntou/modev-backend/gen"
	"net/http"
)

type TweetController struct {
	usecase.TweetUseCase
}

func NewTweetController(t usecase.TweetUseCase) *TweetController {
	return &TweetController{
		t,
	}
}

func (c *TweetController) GetTweets(w http.ResponseWriter, r *http.Request, p gen.GetTweetsParams) {
	ctx := context.Background()
	tt, err := c.TweetUseCase.FindByCategory(ctx, p.Category)
	if err != nil {
		return
	}

	RespondOK(w, tt)
}
