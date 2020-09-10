package pkg

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

// Privacy of the video
type Privacy int

// The privacy options : public , unlisted ,  private
const (
	Public Privacy = iota
	Unlisted
	Private
)

func (p Privacy) String() string {
	return [...]string{"public", "unlisted", "private"}[p]
}

// Video contains all the fields we need for uploading a video
type Video struct {
	Path, Title, Description, Category, Keywords, FileName string
	Privacy                                                Privacy
}

//ListVideos list all the videos we need to upload
func ListVideos(dirPath string, lessons []Lesson) []Video {
	list := []Video{}
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Fatal(fmt.Errorf("the path: %s was not found", dirPath))
	}

	for _, file := range files {
		filePath := dirPath + `\` + file.Name()
		privacy := Unlisted
		video := Video{
			Path: filePath, Title: setVideoName(file.Name(), lessons),
			Description: "Hi", Category: "22", Keywords: "", FileName: file.Name(), Privacy: privacy}
		list = append(list, video)

	}

	return list
}

func setVideoName(fileName string, lessons []Lesson) string {

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

// VideoUploadedToYoutube deletes videos that uploaded successfully
func VideoUploadedToYoutube(vid Video) {

	if e := os.Remove(vid.Path); e != nil {
		log.Fatal(e)
	}
	println("%S deleted successfully ", vid.Title)

}
