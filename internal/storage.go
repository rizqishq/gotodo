package internal

import (
	"encoding/json"
	"errors"
	"os"
)

const fileName = "tasks.json"

func LoadTasks() ([]Task, error) {
	var tasks []Task

	data, err := os.ReadFile(fileName)
	if errors.Is(err, os.ErrNotExist) {
		return []Task{}, nil
	}
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func SaveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0644)
}
