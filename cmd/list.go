package cmd

import (
	"fmt"

	todo "github.com/rizqishq/gotodo/internal"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list [todo|done|in-progress]",
	Short: "List all tasks or by status",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := todo.LoadTasks()
		if err != nil {
			fmt.Println("❌ Error loading tasks:", err)
			return
		}

		var filter string
		if len(args) > 0 {
			filter = args[0]
		}

		if len(tasks) == 0 {
			fmt.Println("📭 No tasks found.")
			return
		}

		fmt.Printf("📋 Listing Tasks (total: %d)\n\n", len(tasks))
		for _, t := range tasks {
			if filter == "" || t.Status == filter {
				todo.PrintTaskVerbose(t)
				fmt.Println("------------------------------------")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
