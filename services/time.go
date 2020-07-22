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
