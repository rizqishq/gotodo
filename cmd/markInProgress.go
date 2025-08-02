package cmd

import (
	"fmt"
	"strconv"
	"time"

	todo "github.com/rizqishq/gotodo/internal"
	"github.com/spf13/cobra"
)

var markInProgressCmd = &cobra.Command{
	Use:   "mark-in-progress <id>",
	Short: "Mark a task as in-progress",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("❌ Invalid ID")
			return
		}

		tasks, err := todo.LoadTasks()
		if err != nil {
			fmt.Println("❌ Error loading tasks:", err)
			return
		}

		index, err := todo.FindTaskIndex(tasks, id)
		if err != nil {
			fmt.Println("❌ Task not found")
			return
		}

		tasks[index].Status = todo.InProgress
		tasks[index].UpdatedAt = time.Now()

		if err := todo.SaveTasks(tasks); err != nil {
			fmt.Println("❌ Error saving task:", err)
			return
		}

		fmt.Println("🚧 Task marked as in-progress.")
		todo.PrintTaskVerbose(tasks[index])
	},
}

func init() {
	rootCmd.AddCommand(markInProgressCmd)
}
