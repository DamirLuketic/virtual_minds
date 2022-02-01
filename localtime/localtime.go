package localtime

import (
	"time"
)

func (t *TimeImpl) CurrentDateWithHour() (*time.Time, error) {
	tn := t.now()
	timeParsed := time.Date(tn.Year(), tn.Month(), tn.Day(), tn.Hour(), 0, 0, 0, tn.Location())
	return &timeParsed, nil
}
