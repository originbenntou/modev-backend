package usecase

import (
	"context"
	"github.com/originbenntou/modev-backend/application/service"
	"github.com/originbenntou/modev-backend/domain/model"
	"github.com/originbenntou/modev-backend/gen"
)

type TweetUseCase interface {
	FindByCategory(ctx context.Context, category gen.GetTweetsParamsCategory) ([]*model.Tweet, error)
}

type tweetUseCase struct {
	service.TweetService
}

func NewTweetUseCase(t service.TweetService) TweetUseCase {
	return &tweetUseCase{
		t,
	}
}

func (u *tweetUseCase) FindByCategory(ctx context.Context, category gen.GetTweetsParamsCategory) ([]*model.Tweet, error) {
	tweets, err := u.TweetService.FindByCategory(ctx, category)
	if err != nil {
		return nil, err
	}
	return tweets, nil
}
