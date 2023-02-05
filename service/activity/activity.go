package activity

import "time"

type Activity struct {
	ActivityID string
	Name       string
	Data       string
	Time       time.Time
}
