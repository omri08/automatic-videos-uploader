package db

import (
	"uploader/models"
)

//LessonService is the interface to make the Lesson Model intercate with a database
type LessonService interface {
	InitDB() error
	AddLesson(l models.Lesson) error
	LoadLessons() []models.Lesson
	DeleteLesson(name string) error
}