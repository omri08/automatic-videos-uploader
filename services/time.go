package services

import (
	"strconv"
	"strings"
	"time"
)

//DateToDay receving string of a specific date, and returns the day of this specific date
func DateToDay(s string) string {
	ANSIC := "Mon Jan _2 15:04:05 2006"
	temp := strings.Split(s, "-")
	year, _ := strconv.Atoi(temp[0])
	monthNumric, _ := strconv.Atoi(temp[1])
	month := time.Month(monthNumric)
	day, _ := strconv.Atoi(temp[2])
	t := time.Date(year, month, day, 10, 23, 0, 0, time.UTC)

	return strings.Split(t.Format(ANSIC), " ")[0]
}

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
