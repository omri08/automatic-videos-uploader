package services

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"
	"youtube/models"
)

//ListVideos list all the videos we need to upload
func ListVideos(dirPath string, lessons []models.Lesson) []models.Video {
	list := []models.Video{}
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Fatal(fmt.Errorf("the path: %s was not found", dirPath))
	}

	for _, file := range files {
		filePath := dirPath + `\` + file.Name()
		privacy := models.Unlisted
		video := models.Video{
			Path: filePath, Title: setVideoName(file.Name(), lessons),
			Description: "Hi", Category: "22", Keywords: "", FileName: file.Name(), Privacy: privacy}
		list = append(list, video)

	}

	return list
}

func setVideoName(fileName string, lessons []models.Lesson) string {

	timeRecorded := StringToTime(fileName)
	for _, lesson := range lessons {

		if int(timeRecorded.Weekday()) == lesson.Day {
			start := time.Date(timeRecorded.Year(), timeRecorded.Month(), timeRecorded.Day(), lesson.Starts-1, 55, 0, 0, timeRecorded.Location())
			end := time.Date(timeRecorded.Year(), timeRecorded.Month(), timeRecorded.Day(), lesson.Ends, 0, 0, 0, timeRecorded.Location())
			var startDiff time.Duration
			if start.After(timeRecorded) {
				startDiff = start.Sub(timeRecorded)
			} else {
				startDiff = timeRecorded.Sub(start)
			}

			if startDiff.Minutes() <= 20 && end.After(timeRecorded) && startDiff.Hours() < 1 {

				return fmt.Sprintf("%s - %d/%d", lesson.Name, timeRecorded.Day(), timeRecorded.Month())
			}
		}
	}

	return fileName
}
