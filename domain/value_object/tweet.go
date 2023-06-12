package value_object

import (
	"fmt"
	"net/url"
	"time"
)

// Category value object
type Category int

const (
	Own Category = iota
	Like
)

func NewCategory(value int) (*Category, error) {
	var category Category
	switch value {
	case int(Own):
		category = Own
	case int(Like):
		category = Like
	default:
		return nil, fmt.Errorf("invalid category value: %d", value)
	}
	return &category, nil
}

func (c Category) Value() int {
	return int(c)
}

func (c Category) String() string {
	var categoryString string
	switch c {
	case Own:
		categoryString = "Own"
	case Like:
		categoryString = "Like"
	default:
		categoryString = "unknown"
	}
	return categoryString
}

type AddDate struct {
	Value time.Time
}

const (
	dateLayout = "2006-01-02T15:04:05Z"
)

func NewAddDate(value string) (*AddDate, error) {
	addDate, err := time.Parse(dateLayout, value)
	if err != nil {
		return nil, fmt.Errorf("invalid add_date value: %s", value)
	}
	return &AddDate{Value: addDate}, nil
}

func (a AddDate) String() string {
	return a.Value.Format(dateLayout)
}

// URL value object
type URL struct {
	Value string
}

func NewURL(urlString string) (*URL, error) {
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, fmt.Errorf("invalid URL")
	}

	return &URL{Value: urlString}, nil
}

// Tag value object
type Tag struct {
	Value string
}
