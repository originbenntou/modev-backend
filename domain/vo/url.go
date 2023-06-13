package vo

import (
	"fmt"
	"net/url"
)

// URL value object
type URL struct {
	Value string
}

func NewURL(urlString string) (*URL, error) {
	u, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, fmt.Errorf("invalid URL: %s", u)
	}

	return &URL{Value: urlString}, nil
}

func (u *URL) Equal(other *URL) bool {
	return u.Value == other.Value
}
