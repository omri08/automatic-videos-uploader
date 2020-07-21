package models

import "log"

// Privacy of the video
type Privacy int

// The privacy options : 1 = public , 2 = unlisted , 3 = private
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
	Path        string
	Title       string
	Description string
	Category    string
	Keywords    string
	Privacy     Privacy
}

//NewVideo is a  constructor of Video
func NewVideo(path, title, description, category, keywords string, privacy Privacy) *Video {
	video := Video{}
	if path == "" {
		log.Fatalf("You must provide a path of a video file to upload")
	}

	return &video
}
