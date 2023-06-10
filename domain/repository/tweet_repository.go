package repository

import (
	"context"
	"github.com/originbenntou/modev-backend/domain/model"
)

type TweetRepository interface {
	FindByCategory(ctx context.Context, category string) (*model.Tweet, error)
}
