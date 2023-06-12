package database

import (
	"context"
	"fmt"
	"github.com/originbenntou/modev-backend/adapter/mysql"
	"github.com/originbenntou/modev-backend/domain/model"
	"github.com/originbenntou/modev-backend/domain/repository"
	vo "github.com/originbenntou/modev-backend/domain/value_object"
)

const TableName = "TWEET"

type tweetRepository struct {
	db *mysql.DB
}

func NewTweetRepository(db *mysql.DB) repository.TweetRepository {
	return &tweetRepository{db}
}

func (r *tweetRepository) FindByCategory(ctx context.Context, category *vo.Category) ([]*model.TweetModel, error) {
	q := fmt.Sprintf(`
		select
			id, category, add_date, url, created_at, updated_at
		from
		    %s
		where
		    category = :category
	`, TableName)

	rows, err := r.db.NamedQueryContext(ctx, q, map[string]interface{}{
		"category": category.Value(),
	})
	if err != nil {
		return nil, err
	}

	results := make([]*model.TweetModel, 0)
	for rows.Next() {
		var tweet model.TweetModel
		if err := rows.StructScan(&tweet); err != nil {
			return nil, err
		}
		results = append(results, &tweet)
	}

	return results, nil
}
