package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("📦 Task Tracker CLI")
		fmt.Println("Usage:")
		fmt.Println("  add <description>")
		fmt.Println("  list [todo|done|in-progress]")
		fmt.Println("  update <id> <new description>")
		fmt.Println("  delete <id>")
		fmt.Println("  mark-in-progress <id>")
		fmt.Println("  mark-done <id>")
		return
	}

	command := os.Args[1]
	tasks, err := LoadTasks()
	if err != nil {
		fmt.Println("❌ Error loading tasks:", err)
		return
	}

	switch command {

	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: add <task description>")
			return
		}

		desc := strings.Join(os.Args[2:], " ")
		task := Task{
			ID:          GetNextID(tasks),
			Description: desc,
			Status:      Todo,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
		tasks = append(tasks, task)

		if err := SaveTasks(tasks); err != nil {
			fmt.Println("❌ Failed to save task:", err)
			return
		}

		fmt.Println("✅ Task added successfully!")
		printTaskVerbose(task)

	case "list":
		var filter string
		if len(os.Args) >= 3 {
			filter = os.Args[2]
		}

		if len(tasks) == 0 {
			fmt.Println("📭 No tasks found.")
			return
		}

		fmt.Printf("📋 Listing Tasks (total: %d)\n\n", len(tasks))
		for _, t := range tasks {
			if filter == "" || string(t.Status) == filter {
				printTaskVerbose(t)
				fmt.Println("------------------------------------")
			}
		}

	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Usage: update <id> <new description>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("❌ Invalid ID:", os.Args[2])
			return
		}
		newDesc := strings.Join(os.Args[3:], " ")
		index, err := FindTaskIndex(tasks, id)
		if err != nil {
			fmt.Println("❌ Task not found")
			return
		}
		tasks[index].Description = newDesc
		tasks[index].UpdatedAt = time.Now()

		if err := SaveTasks(tasks); err != nil {
			fmt.Println("❌ Error saving task:", err)
			return
		}
		fmt.Println("✏️ Task updated successfully!")
		printTaskVerbose(tasks[index])

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: delete <id>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("❌ Invalid ID:", os.Args[2])
			return
		}
		index, err := FindTaskIndex(tasks, id)
		if err != nil {
			fmt.Println("❌ Task not found")
			return
		}
		deleted := tasks[index]
		tasks = append(tasks[:index], tasks[index+1:]...)

		if err := SaveTasks(tasks); err != nil {
			fmt.Println("❌ Error saving task:", err)
			return
		}
		fmt.Println("🗑️ Task deleted successfully!")
		fmt.Printf("🆔 ID         : %d\n", deleted.ID)
		fmt.Printf("📄 Description: %s\n", deleted.Description)

	case "mark-in-progress":
		if len(os.Args) < 3 {
			fmt.Println("Usage: mark-in-progress <id>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("❌ Invalid ID")
			return
		}
		index, err := FindTaskIndex(tasks, id)
		if err != nil {
			fmt.Println("❌ Task not found")
			return
		}
		tasks[index].Status = InProgress
		tasks[index].UpdatedAt = time.Now()
		if err := SaveTasks(tasks); err != nil {
			fmt.Println("❌ Error saving task:", err)
			return
		}
		fmt.Println("🚧 Task marked as in-progress.")
		printTaskVerbose(tasks[index])

	case "mark-done":
		if len(os.Args) < 3 {
			fmt.Println("Usage: mark-done <id>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("❌ Invalid ID")
			return
		}
		index, err := FindTaskIndex(tasks, id)
		if err != nil {
			fmt.Println("❌ Task not found")
			return
		}
		tasks[index].Status = Done
		tasks[index].UpdatedAt = time.Now()
		if err := SaveTasks(tasks); err != nil {
			fmt.Println("❌ Error saving task:", err)
			return
		}
		fmt.Println("✅ Task marked as done.")
		printTaskVerbose(tasks[index])

	default:
		fmt.Println("❓ Unknown command:", command)
		fmt.Println("Available commands: add, list, update, delete, mark-in-progress, mark-done")
	}
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

func printTaskVerbose(t Task) {
	fmt.Printf("🆔 ID         : %d\n", t.ID)
	fmt.Printf("📄 Description: %s\n", t.Description)
	fmt.Printf("📌 Status     : %s\n", formatStatus(t.Status))
	fmt.Printf("🕒 Created At : %s\n", formatTime(t.CreatedAt))
	fmt.Printf("🕒 Updated At : %s\n", formatTime(t.UpdatedAt))
}
