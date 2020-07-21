package main

import (
	"fmt"
	"youtube/services"
)

func main() {
	fmt.Println("Hello World")
	list := services.ListVideos(`C:\Users\Omri\Desktop\records`)
	services.UploadToYoutube(list)

}
