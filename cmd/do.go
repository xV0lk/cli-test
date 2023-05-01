package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/xV0lk/cli-test/db"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as complete",
	Run: func(cmd *cobra.Command, args []string) {

		var ids []int
		ctIds := currentTasksIds()

		for _, taskNum := range args {
			id, err := strconv.Atoi(taskNum)
			if err != nil {
				fmt.Printf("Failed to parse argument: %s\n", taskNum)
				os.Exit(1)
			} else {
				ids = append(ids, id)
			}
		}
		for _, id := range ids {
			if id <= 0 || id > len(ctIds) {
				fmt.Printf("Invalid task number: %d\n", id)
				continue
			}
			if Contains(ctIds, id) {
				db.DeleteTask(id)
				fmt.Printf("Deleted task with id: %d\n", id)
			} else {
				fmt.Printf("You don't have a task with id: %d\n", id)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)

}

func currentTasksIds() []int {
	tasks, err := db.AllTasks()
	var tIds []int
	if err != nil {
		fmt.Println("Error: ", err)
		return tIds
	}
	for _, task := range tasks {
		tIds = append(tIds, task.Key)
	}
	return tIds
}

func Contains[T comparable](sli []T, v T) bool {
	for _, item := range sli {
		if v == item {
			return true
		}
	}
	return false
}
