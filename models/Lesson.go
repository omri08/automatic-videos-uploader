package models

//Lesson contains the info of each leacture we have, so we will be able to set the right values for the videos
type Lesson struct {
	Name   string `json:"name"`
	Day    int    `json:"day"`
	Starts int    `json:"starts"`
	Ends   int    `json:"ends"`
}
