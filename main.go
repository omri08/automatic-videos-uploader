package main

import (
	"io/ioutil"
	"os"
	"youtube/services"
)

func main() {
	jsonFile, _ := os.Open("./data/schedule.json")
	defer jsonFile.Close()
	data, _ := ioutil.ReadAll(jsonFile)
	lessons := services.JSONToLessonsArr(data)
	list := services.ListVideos(`C:\Users\Omri\Desktop\records`, lessons)
	services.UploadToYoutube(list)

}
