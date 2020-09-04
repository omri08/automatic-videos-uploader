package services

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"uploader/models"
)

var jsonPath = `C:\Users\Omri\go\src\youtube\data\schedule.json`

//LoadLessons converts json data into lesson Array
func LoadLessons() []models.Lesson {
	jsonFile, err := os.Open(jsonPath)
	defer jsonFile.Close()
	if err != nil {
		return nil
	}

	data, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil
	}
	var lessons []models.Lesson
	json.Unmarshal([]byte(data), &lessons)

	return lessons
}

//AddLesson Addes a lessons to schedule.json
func AddLesson(les models.Lesson) error {

	jsonFile, err := os.OpenFile(jsonPath, os.O_RDWR, 777)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	data, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	var lessons []models.Lesson
	json.Unmarshal([]byte(data), &lessons)
	lessons = append(lessons, les)

	jsonArr, err := json.MarshalIndent(lessons, "", "\t")
	if err != nil {
		return err
	}

	jsonFile.Seek(0, os.SEEK_SET)

	if _, err := jsonFile.WriteString(string(jsonArr)); err != nil {
		return err
	}

	return nil
}
