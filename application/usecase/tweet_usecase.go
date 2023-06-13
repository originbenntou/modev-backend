package usecase

import (
	"context"
	"github.com/originbenntou/modev-backend/domain/entity"
	"github.com/originbenntou/modev-backend/domain/service"
	vo "github.com/originbenntou/modev-backend/domain/vo"
)

type TweetUseCase interface {
	FindByCategory(ctx context.Context, category *vo.Category) ([]*entity.TweetEntity, error)
}

type tweetUseCase struct {
	service.TweetService
}

func NewTweetUseCase(t service.TweetService) TweetUseCase {
	return &tweetUseCase{
		t,
	}
}

func (u *tweetUseCase) FindByCategory(ctx context.Context, category *vo.Category) ([]*entity.TweetEntity, error) {
	tweets, err := u.TweetService.FindByCategory(ctx, category)
	if err != nil {
		return nil, err
	}
	return tweets, nil
}
