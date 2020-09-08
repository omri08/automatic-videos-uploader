package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"uploader/config"
	"uploader/models"
)

//GlobalMySQL the only variable to interact with MySQL database
var GlobalMySQL = &LessonServiceSQL{}

func init() {
	GlobalMySQL.InitDB()
}

// LessonServiceSQL implements LessonService with MySQL connection
type LessonServiceSQL struct {
	DB *sql.DB
}

func (s *LessonServiceSQL) connectToDb() error {
	db, err := sql.Open("mysql", config.UserName+":"+config.Password+"@tcp(127.0.0.1:3306)/")

	if err != nil {
		return err
	}
	println("Connected To DataBase")
	s.DB = db

	return nil

}

//InitDB  connection to the database, REMEMBER TO CLOSE THE DATABASE!!!
func (s *LessonServiceSQL) InitDB() error {
	if err := s.connectToDb(); err != nil {
		return err
	}

	if _, err := s.DB.Exec("CREATE DATABASE IF NOT EXISTS UploaderDB"); err != nil {
		println("Failed to created DATABASE")
		return err
	}

	if _, err := s.DB.Exec("USE UploaderDB"); err != nil {
		println("Failed using")
		return err
	}

	if _, err := s.DB.Exec(
		"CREATE TABLE IF NOT EXISTS Lesson(name varchar(255) PRIMARY KEY NOT NULL,day INT NOT NULL,starts INT NOT NULL,ends INT NOT NULL)"); err != nil {
		println("Failed to create lesson table")
		return err
	}

	return nil
}

//AddLesson adds a lesson to Lesson table
func (s *LessonServiceSQL) AddLesson(l models.Lesson) error {
	stmt, err := s.DB.Prepare("INSERT INTO Lesson(name,day,starts,ends) VALUES(?,?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(l.Name, l.Day, l.Starts, l.Ends)
	if err != nil {
		return err
	}
	return nil
}

//LoadLessons loads all the lessons in the Lesson table
func (s LessonServiceSQL) LoadLessons() []models.Lesson {
	rows, err := s.DB.Query("SELECT * from Lesson")
	if err != nil {
		return nil
	}
	var lessons []models.Lesson
	defer rows.Close()
	for rows.Next() {
		var l models.Lesson
		err := rows.Scan(&l.Name, &l.Day, &l.Starts, &l.Ends)
		if err != nil {
			return nil
		}
		lessons = append(lessons, l)
	}

	return lessons
}

//DeleteLesson delets a lesson from the lesson table
func (s LessonServiceSQL) DeleteLesson(name string) error {
	stmt, err := s.DB.Prepare("DELETE from Lesson where name =?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(name)
	if err != nil {
		return err
	}

	return nil
}
