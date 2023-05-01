package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/xV0lk/cli-test/db"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to the task manager",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		db.CreateTask(task)
		fmt.Printf("Added \"%s\" to your tasks\n", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
