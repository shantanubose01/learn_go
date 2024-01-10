package handler

import (
	"encoding/json"
	"os"
)

type Todo struct {
	ID     int    `json : "_id"`
	Task   string `json : "task"`
	Status string `json : "status"`
}

func readTodosFromFile(filePath string) ([]Todo, error) {
	var todos []Todo
	file, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(file, &todos)
	return todos, err
}

func writeTodosToFile(filePath string, todo []Todo) error {
	data, err := json.MarshalIndent(todo, "", "")
	if err != nil {
		panic(err)
	}

	return os.WriteFile(filePath, data, os.ModePerm)

}

func GetAllTodosHandler() []Todo {

	data, err := readTodosFromFile(("data.json"))

	if err != nil {
		panic(err)
	}

	return data
}
