package repository

import (
	"context"
	"github.com/originbenntou/modev-backend/domain/model"
	vo "github.com/originbenntou/modev-backend/domain/value_object"
)

type TweetRepository interface {
	FindByCategory(ctx context.Context, category *vo.Category) ([]*model.TweetModel, error)
}
