package service

import (
	"context"
	"fmt"
	"github.com/originbenntou/modev-backend/domain/model"
	"github.com/originbenntou/modev-backend/domain/repository"
)

type TweetService interface {
	FindByCategory(ctx context.Context, category string) (*model.Tweet, error)
}

type tweetService struct {
	repository.TweetRepository
}

func NewTweetService(t repository.TweetRepository) TweetService {
	return &tweetService{
		t,
	}
}

func (s *tweetService) FindByCategory(ctx context.Context, category string) (*model.Tweet, error) {
	tt, err := s.TweetRepository.FindByCategory(ctx, category)
	if err != nil {
		return nil, err
	}

	fmt.Println(tt)

	return tt, nil
}
