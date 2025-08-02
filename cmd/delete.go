package cmd

import (
	"fmt"
	"strconv"

	todo "github.com/rizqishq/gotodo/internal"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "Delete a task by ID",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("❌ Invalid ID:", args[0])
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

		deleted := tasks[index]
		tasks = append(tasks[:index], tasks[index+1:]...)

		if err := todo.SaveTasks(tasks); err != nil {
			fmt.Println("❌ Error saving task:", err)
			return
		}

		fmt.Println("🗑️ Task deleted successfully!")
		fmt.Printf("🆔 ID         : %d\n", deleted.ID)
		fmt.Printf("📄 Description: %s\n", deleted.Description)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
