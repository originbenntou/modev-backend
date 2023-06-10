package usecase

import (
	"context"
	"github.com/originbenntou/modev-backend/application/service"
	"github.com/originbenntou/modev-backend/domain/model"
)

type TweetUseCase interface {
	FindByCategory(ctx context.Context, category string) (*model.Tweet, error)
}

type tweetUseCase struct {
	service.TweetService
}

func NewTweetUseCase(t service.TweetService) TweetUseCase {
	return &tweetUseCase{
		t,
	}
}

func (u *tweetUseCase) FindByCategory(ctx context.Context, category string) (*model.Tweet, error) {
	tt, err := u.TweetService.FindByCategory(ctx, category)
	if err != nil {
		return nil, err
	}
	return tt, nil
}
