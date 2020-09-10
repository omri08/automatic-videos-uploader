package pkg

import (
	"fmt"
	"testing"
	"time"
)

//TestStringToTime test the StringToTime Function
func TestStringToTime(t *testing.T) {
	var tests = []struct {
		s    string
		want time.Time
	}{{"2020-07-22 09-03-25", time.Date(2020, time.July, 22, 9, 3, 25, 0, time.UTC)}}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s %v \n", tt.s, tt.want)
		t.Run(testname, func(t *testing.T) {
			ans := StringToTime2(tt.s)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}
