package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	todo "github.com/rizqishq/gotodo/internal"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update <id> <new description>",
	Short: "Update a task's description",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("❌ Invalid ID:", args[0])
			return
		}

		newDesc := strings.Join(args[1:], " ")

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

		tasks[index].Description = newDesc
		tasks[index].UpdatedAt = time.Now()

		if err := todo.SaveTasks(tasks); err != nil {
			fmt.Println("❌ Error saving task:", err)
			return
		}

		fmt.Println("✏️ Task updated successfully!")
		todo.PrintTaskVerbose(tasks[index])
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
