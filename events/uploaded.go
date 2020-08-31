package events

import (
	"log"
	"os"
	"youtube/models"
)

// VideosUploadedToYouTube deletes videos that uploaded successfully
func VideosUploadedToYouTube(vid models.Video) {

	if e := os.Remove(vid.Path); e != nil {
		log.Fatal(e)
	}
	println("%S deleted successfully ", vid.Title)

}
