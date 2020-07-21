package services

import (
	"fmt"
	"io/ioutil"
	"log"
	"youtube/models"
)

//ListVideos list all the videos we need to upload
func ListVideos(dirPath string) []models.Video {
	list := []models.Video{}
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Fatal(fmt.Errorf("the path: %s was not found", dirPath))
	}

	for _, file := range files {
		filePath := dirPath + `\` + file.Name()
		privacy := models.Private
		video := models.Video{Path: filePath, Title: file.Name(), Description: "Hi", Category: "22", Keywords: "", Privacy: privacy}
		list = append(list, video)

	}

	return list
}
