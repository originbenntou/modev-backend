package database

import (
	"context"
	"github.com/originbenntou/modev-backend/domain/model"
	"github.com/originbenntou/modev-backend/domain/repository"
)

type tweetRepository struct {
}

func NewTweetRepository() repository.TweetRepository {
	return &tweetRepository{}
}

func (r *tweetRepository) FindByCategory(ctx context.Context, category string) (*model.Tweet, error) {
	return &model.Tweet{
		Id: 1,
	}, nil
}
