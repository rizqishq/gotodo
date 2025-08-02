package cmd

import (
	"fmt"
	"strings"
	"time"

	todo "github.com/rizqishq/gotodo/internal"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add <task description>",
	Short: "Add a new task",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := todo.LoadTasks()
		if err != nil {
			fmt.Println("❌ Error loading tasks:", err)
			return
		}

		desc := strings.Join(args, " ")
		task := todo.Task{
			ID:          todo.GetNextID(tasks),
			Description: desc,
			Status:      todo.Todo,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
		tasks = append(tasks, task)

		if err := todo.SaveTasks(tasks); err != nil {
			fmt.Println("❌ Failed to save task:", err)
			return
		}

		fmt.Println("✅ Task added successfully!")
		todo.PrintTaskVerbose(task)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
