package pkg

import (
	"time"
)

// StringToTime take a stirng and returns a valid Time object
func StringToTime(s string) time.Time {
	layOut := "2006-1-2 15-04-05"
	t, _ := time.Parse(layOut, s)
	return t
}
