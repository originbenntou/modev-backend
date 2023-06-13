package model

import "database/sql"

type TweetModel struct {
	Id        uint64 `db:"id"`
	Category  int    `db:"category"`
	AddDate   string `db:"add_date"`
	Url       string `db:"url"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

type TweetWithTagModel struct {
	TweetModel
	Tag sql.NullString `db:"tag_name"`
}
