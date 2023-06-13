package vo

import (
	"fmt"
	"time"
)

type DateTime struct {
	Value time.Time
}

const (
	DatetimeLayout = "2006-01-02T15:04:05Z"
	DateLayout     = "2006-01-02"
)

var (
	loc *time.Location
)

func init() {
	loc, _ = time.LoadLocation("Asia/Tokyo")
}

func NewDateTime(value string) (*DateTime, error) {
	dateTime, err := time.Parse(DatetimeLayout, value)
	if err != nil {
		return nil, fmt.Errorf("invalid dateTime value: %s", value)
	}
	return &DateTime{Value: dateTime.In(loc)}, nil
}

func (d *DateTime) ToString() string {
	return d.Value.Format(DatetimeLayout)
}

func (d *DateTime) ToDateString() string {
	return d.Value.Format(DateLayout)
}

func (d *DateTime) Equal(other *DateTime) bool {
	return d.Value == other.Value
}
