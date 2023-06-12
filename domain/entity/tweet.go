package entity

import (
	vo "github.com/originbenntou/modev-backend/domain/value_object"
	"time"
)

type TweetEntity struct {
	Id        uint64
	Category  vo.Category
	AddDate   vo.AddDate
	Url       vo.URL
	Tags      []vo.Tag
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (t *TweetEntity) ChangeCategory(category vo.Category) {
	t.Category = category
}
