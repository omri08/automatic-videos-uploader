package events

import (
	"log"
	"os"
	"uploader/models"
)

// VideoUploadedToYoutube deletes videos that uploaded successfully
func VideoUploadedToYoutube(vid models.Video) {

	if e := os.Remove(vid.Path); e != nil {
		log.Fatal(e)
	}
	println("%S deleted successfully ", vid.Title)

}
