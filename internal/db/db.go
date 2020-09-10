package db

import (
	"uploader/pkg"
)

//LessonService is the interface to make the Lesson Model intercate with a database
type LessonService interface {
	InitDB() error
	AddLesson(l pkg.Lesson) error
	LoadLessons() []pkg.Lesson
	DeleteLesson(name string) error
	Close()
}

//DB is our way to make the Lesson Model intercate with the DataBase
var DB LessonService = mySQL

func init() {
	DB.InitDB()
}
