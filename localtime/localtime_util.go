package localtime

import "time"

func NewTime() Time {
	return &TimeImpl{}
}

func (t *TimeImpl) now() time.Time {
	return time.Now()
}
