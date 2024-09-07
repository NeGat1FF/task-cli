package cmd

import (
	"fmt"
	"strconv"

	"github.com/NeGat1FF/task-cli/task"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update [task ID]",
	Short: "Update the description of a task by its ID.",
	Long: `Update the description of the task with the specified ID. Provide the task ID and the new description as positional arguments. This command changes the description of the task and updates its 'Updated At' timestamp.

Example usage: task-cli update 1 "New Description"

This will update the task with ID 1 to have the description 'New Description'. If the task with the given ID is found and updated successfully, a confirmation message will be displayed.`,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		val, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		err = task.UpdateTask(val, args[1])
		if err != nil {
			return err
		}

		fmt.Printf("Task with ID: %d updated successfully\n", val)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
