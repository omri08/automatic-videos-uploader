package events

import (
	"log"
	"os"
	"youtube/models"
)

// VideosUploadedToYouTube takes a path to the directory ,list of the videos and deletes the videos that uploaded successfully
func VideosUploadedToYouTube(dirPath string, vidsArr []models.Video) {
	for _, vid := range vidsArr {
		if vid.Uploaded {
			filePath := dirPath + `\` + vid.FileName
			if e := os.Remove(filePath); e != nil {
				log.Fatal(e)
			}
			println("%S deleted successfully ", vid.Title)

		}
	}
}
