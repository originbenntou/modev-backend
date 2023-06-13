package controller

import (
	"context"
	"fmt"
	"github.com/originbenntou/modev-backend/application/usecase"
	vo "github.com/originbenntou/modev-backend/domain/vo"
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

	category, err := exchangeCategory(p.Category)
	if err != nil {
		// FIXME: error respondError
		return
	}

	tweets, err := c.TweetUseCase.FindByCategory(ctx, category)
	if err != nil {
		// FIXME: error respondError
		return
	}

	RespondOK(w, tweets)
}

func exchangeCategory(p gen.GetTweetsParamsCategory) (*vo.Category, error) {
	var c vo.Category

	switch p {
	case gen.Own:
		c = vo.Own
	case gen.Like:
		c = vo.Like
	default:
		return nil, fmt.Errorf("invalid category: %s", p)
	}

	return &c, nil
}
