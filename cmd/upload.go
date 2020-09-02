package cmd

import (
	"uploader/services"
)

func main() {

	lessons := services.JSONToLessonsArr()
	list := services.ListVideos(`C:\Users\Omri\Desktop\records`, lessons)
	services.UploadToYoutube(list)

}
