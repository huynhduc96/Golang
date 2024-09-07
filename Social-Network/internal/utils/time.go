package utils

import (
	"social/internal/constant"
	"time"
)

func StringToTime(dateString string) (time.Time, error) {
	t, err := time.Parse(constant.DOBLayout, dateString)
	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}
