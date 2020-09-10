package pkg

import (
	"strconv"
	"strings"
	"time"
)

// StringToTime take a stirng and returns a valid Time object
func StringToTime(s string) time.Time {
	// timeString[0] is date, timeString[1] is hour
	timeString := strings.Split(s, " ")
	// date[0] = year, date[1] = month, date[2] = day
	date := strings.Split(timeString[0], "-")
	// hourString[0] = hh , hourString[1] = mm , hourString[2] = ss
	hourString := strings.Split(timeString[1], "-")
	monthNumric, _ := strconv.Atoi(date[1])
	month := time.Month(monthNumric)
	day, _ := strconv.Atoi(date[2])
	year, _ := strconv.Atoi(date[0])
	hour, _ := strconv.Atoi(hourString[0])
	min, _ := strconv.Atoi(hourString[1])
	sec, _ := strconv.Atoi(hourString[2])
	t := time.Now()

	validTime := time.Date(year, month, day, hour, min, sec, 0, t.Location())
	return validTime

}
