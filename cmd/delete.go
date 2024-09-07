package cmd

import (
	"fmt"
	"strconv"

	"github.com/NeGat1FF/task-cli/task"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [task ID]",
	Short: "Delete a task by its ID.",
	Long: `"Delete the task specified by its ID. Provide the ID of the task you wish to delete as a positional argument. If the task with the given ID exists, it will be removed from the task list. If successful, a confirmation message with the deleted task's ID will be displayed.

Example usage: task-cli delete 1

"`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		val, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}

		err = task.DeleteTask(val)
		if err != nil {
			return err
		}
		fmt.Printf("Task with ID: %d deleted successfully\n", val)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
