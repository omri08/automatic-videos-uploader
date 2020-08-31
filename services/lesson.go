package services

import (
	"encoding/json"
	"youtube/models"
)

//JSONToLessonsArr converts json data into lesson Array
func JSONToLessonsArr(data []byte) []models.Lesson {
	var lessons []models.Lesson
	json.Unmarshal([]byte(data), &lessons)

	return lessons
}
