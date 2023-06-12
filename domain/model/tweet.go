package model

type TweetModel struct {
	Id        uint64 `db:"id"`
	Category  int    `db:"category"`
	AddDate   string `db:"add_date"`
	Url       string `db:"url"`
	Tags      []string
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}
