package internal

import (
	"errors"
	"fmt"
	"time"
)

func GetNextID(tasks []Task) int {
	max := 0
	for _, t := range tasks {
		if t.ID > max {
			max = t.ID
		}
	}
	return max + 1
}

func FindTaskIndex(tasks []Task, id int) (int, error) {
	for i, t := range tasks {
		if t.ID == id {
			return i, nil
		}
	}
	return -1, errors.New("task not found")
}

func formatStatus(s string) string {
	switch s {
	case Todo:
		return "todo 📝"
	case InProgress:
		return "in-progress 🚧"
	case Done:
		return "done ✅"
	default:
		return string(s)
	}
}

func formatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04")
}

func PrintTaskVerbose(t Task) {
	fmt.Printf("🆔 ID         : %d\n", t.ID)
	fmt.Printf("📄 Description: %s\n", t.Description)
	fmt.Printf("📌 Status     : %s\n", formatStatus(t.Status))
	fmt.Printf("🕒 Created At : %s\n", formatTime(t.CreatedAt))
	fmt.Printf("🕒 Updated At : %s\n", formatTime(t.UpdatedAt))
}
