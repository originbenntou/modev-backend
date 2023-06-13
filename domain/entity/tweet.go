package entity

import (
	vo "github.com/originbenntou/modev-backend/domain/vo"
)

type TweetEntity struct {
	Id        uint64
	Category  vo.Category
	AddDate   string
	Url       vo.URL
	Tags      []vo.Tag
	CreatedAt vo.DateTime
	UpdatedAt vo.DateTime
}

func (t *TweetEntity) ChangeCategory(category vo.Category) {
	t.Category = category
}
