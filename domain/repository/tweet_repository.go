package repository

import (
	"context"
	"github.com/originbenntou/modev-backend/domain/model"
	vo "github.com/originbenntou/modev-backend/domain/vo"
)

type TweetRepository interface {
	FindByCategory(ctx context.Context, category *vo.Category) ([]*model.TweetWithTagModel, error)
}
