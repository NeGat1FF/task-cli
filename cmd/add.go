package cmd

import (
	"fmt"

	"github.com/NeGat1FF/task-cli/task"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:                   "add [task description]",
	Short:                 "Create a new task with the given description.",
	Long:                  "Create a new task with the provided description. This command adds a new task to the task list with an automatically assigned ID, and sets its status to 'TODO'. The task description is required as a positional argument. Upon successful addition, the ID of the newly created task will be displayed.",
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := task.AddTask(args[0])
		if err != nil {
			return err
		}

		fmt.Printf("Task added successfully (ID: %d)\n", id)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
