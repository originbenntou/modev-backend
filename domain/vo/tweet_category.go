package vo

import (
	"fmt"
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

func (c *Category) Value() int {
	return int(*c)
}

func (c *Category) ToString() string {
	var categoryString string
	switch *c {
	case Own:
		categoryString = "Own"
	case Like:
		categoryString = "Like"
	default:
		categoryString = "unknown"
	}
	return categoryString
}

func (c *Category) Equal(other *Category) bool {
	return c.Value() == other.Value()
}
