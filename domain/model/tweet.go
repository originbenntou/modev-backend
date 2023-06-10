package model

type Tweet struct {
	Id      uint64 `db:"id"`
	addDate string `db:"name"`
	url     string `db:"url"`
	tags    []string
}
