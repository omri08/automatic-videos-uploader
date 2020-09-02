package services

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"uploader/models"
)

var jsonPath = `C:\Users\Omri\go\src\youtube\data\schedule.json`

//JSONToLessonsArr converts json data into lesson Array
func JSONToLessonsArr() []models.Lesson {
	jsonFile, _ := os.Open(jsonPath)
	defer jsonFile.Close()
	data, _ := ioutil.ReadAll(jsonFile)
	var lessons []models.Lesson
	json.Unmarshal([]byte(data), &lessons)

	return lessons
}

//AddLessonToJSON Addes a lessons to schedule.json
func AddLessonToJSON(les models.Lesson) error {

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
