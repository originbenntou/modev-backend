package database

import (
	"context"
	"fmt"
	"github.com/originbenntou/modev-backend/adapter/mysql"
	"github.com/originbenntou/modev-backend/domain/model"
	"github.com/originbenntou/modev-backend/domain/repository"
	"github.com/originbenntou/modev-backend/gen"
)

const TABLE_NAME = "TWEET"

type tweetRepository struct {
	db *mysql.DB
}

func NewTweetRepository(db *mysql.DB) repository.TweetRepository {
	return &tweetRepository{db}
}

func (r *tweetRepository) FindByCategory(ctx context.Context, category gen.GetTweetsParamsCategory) ([]*model.Tweet, error) {
	q := fmt.Sprintf(`
		select
			id, category, add_date, url, created_at, updated_at
		from
		    %s
		where
		    category = :category
	`, TABLE_NAME)

	rows, err := r.db.NamedQueryContext(ctx, q, map[string]interface{}{
		"category": "1", // FIXME
	})
	if err != nil {
		return nil, err
	}

	results := make([]*model.Tweet, 0)
	for rows.Next() {
		var tweet model.Tweet
		if err := rows.StructScan(&tweet); err != nil {
			return nil, err
		}
		results = append(results, &tweet)
	}

	return results, nil
}
