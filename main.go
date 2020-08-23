package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"youtube/models"
	"youtube/services"
)

func main() {
	jsonFile, _ := os.Open("./data/schedule.json")
	data, _ := ioutil.ReadAll(jsonFile)
	var lessons []models.Lesson
	json.Unmarshal([]byte(data), &lessons)
	list := services.ListVideos(`C:\Users\Omri\Desktop\records`, lessons)
	fmt.Println(list)

}
