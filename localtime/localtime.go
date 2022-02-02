package localtime

import (
	"time"
)

const (
	DateLayout = "2006-01-02"
)

func (t *TimeImpl) CurrentDateWithHour() (*time.Time, error) {
	tn := t.now()
	timeParsed := time.Date(tn.Year(), tn.Month(), tn.Day(), tn.Hour(), 0, 0, 0, tn.Location())
	return &timeParsed, nil
}

func (t *TimeImpl) ValidateDate(date string) (string, error) {
	ts, err := time.Parse(DateLayout, date)
	if err != nil {
		return "", err
	}
	return ts.Format(DateLayout), nil
}
