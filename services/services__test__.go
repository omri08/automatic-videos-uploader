package services

import (
	"fmt"
	"testing"
	"youtube/models"
)

//TestSetVideoName tests the TestSetVideoName function
func TestSetVideoName(t *testing.T) {
	var tests = []struct {
		fileName string
		lessons  []models.Lesson
		want     string
	}{
		{"2020-07-22 09-03-25", []models.Lesson{{Name: "IntroTo", Day: 3, Starts: 9, Ends: 11}, {Name: "Par", Day: 1, Starts: 13, Ends: 17}}, "IntroTo - 22/7"},
		{"2020-07-22 09-03-25", []models.Lesson{{Name: "Par", Day: 3, Starts: 13, Ends: 17}, {Name: "IntroTo", Day: 3, Starts: 9, Ends: 11}}, "IntroTo - 22/7"},
		{"2020-07-22 09-03-25", []models.Lesson{{}}, "2020-07-22 09-03-25"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s %v %s", tt.fileName, tt.lessons, tt.want)
		t.Run(testname, func(t *testing.T) {
			ans := setVideoName(tt.fileName, tt.lessons)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}
