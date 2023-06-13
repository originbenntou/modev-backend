package database

import (
	"context"
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/originbenntou/modev-backend/adapter/mysql"
	"github.com/originbenntou/modev-backend/domain/model"
	"github.com/originbenntou/modev-backend/domain/repository"
	"github.com/originbenntou/modev-backend/domain/vo"
)

const TweetTable = "TWEET"
const TweetTagTable = "TWEET_TAG"
const TagTable = "TAG"

type tweetRepository struct {
	db *mysql.DB
}

func NewTweetRepository(db *mysql.DB) repository.TweetRepository {
	return &tweetRepository{db}
}

func (r *tweetRepository) FindByCategory(ctx context.Context, category *vo.Category) ([]*model.TweetWithTagModel, error) {
	// TODO: tags を array_agg するパターンも書いてみる
	q := fmt.Sprintf(`
		select
			t1.id,
			t1.category,
			t1.add_date,
			t1.url, t3.name as tag_name,
			t1.created_at,
			t1.updated_at
		from
		    %s t1
			left join %s t2
				on t1.id = t2.tweet_id 
			left join %s t3
				on t2.tag_id = t3.id
		 where
		    category = :category
	`, TweetTable, TweetTagTable, TagTable)

	rows, err := r.db.NamedQueryContext(ctx, q, map[string]any{
		"category": category.Value(),
	})
	if err != nil {
		log.Warn(err)
		return nil, err
	}

	results := make([]*model.TweetWithTagModel, 0)
	for rows.Next() {
		var tweet model.TweetWithTagModel
		if err := rows.StructScan(&tweet); err != nil {
			log.Warn(err)
			return nil, err
		}
		results = append(results, &tweet)
	}

	return results, nil
}
