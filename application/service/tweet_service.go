package service

import (
	"context"
	"github.com/originbenntou/modev-backend/domain/model"
	"github.com/originbenntou/modev-backend/domain/repository"
	"github.com/originbenntou/modev-backend/gen"
)

type TweetService interface {
	FindByCategory(ctx context.Context, category gen.GetTweetsParamsCategory) ([]*model.Tweet, error)
}

type tweetService struct {
	repository.TweetRepository
}

func NewTweetService(t repository.TweetRepository) TweetService {
	return &tweetService{
		t,
	}
}

func (s *tweetService) FindByCategory(ctx context.Context, category gen.GetTweetsParamsCategory) ([]*model.Tweet, error) {
	tweets, err := s.TweetRepository.FindByCategory(ctx, category)
	if err != nil {
		return nil, err
	}

	return tweets, nil
}
