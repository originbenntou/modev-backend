package repository

import (
	"context"
	"github.com/originbenntou/modev-backend/domain/model"
	"github.com/originbenntou/modev-backend/gen"
)

type TweetRepository interface {
	FindByCategory(ctx context.Context, category gen.GetTweetsParamsCategory) ([]*model.Tweet, error)
}
