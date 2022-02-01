package localtime

import "time"

type TimeImpl struct {
}

type Time interface {
	CurrentDateWithHour() (*time.Time, error)
}
