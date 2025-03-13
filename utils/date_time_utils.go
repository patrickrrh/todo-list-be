package utils

import (
	"log"
	"time"
)

func ParseDueTime(dateStr, timeStr string) (time.Time, error) {
	if dateStr == "" {
		return time.Time{}, nil
	}

	if timeStr == "" {
		timeStr = "23:59:00"
	}

	dueTimeStr := dateStr + " " + timeStr
	layout := "2006-01-02 15:04:05"

	dueTime, err := time.ParseInLocation(layout, dueTimeStr, time.Local)
	if err != nil {
		log.Println("[utils][date_utils.go][ParseDueTime] Error parsing due date and time:", err.Error())
		return time.Time{}, err
	}

	return dueTime.UTC(), nil
}

func IsOverdue(dueTime time.Time) bool {
	return !dueTime.IsZero() && dueTime.Before(time.Now().UTC())
}
