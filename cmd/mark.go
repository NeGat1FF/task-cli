package cmd

import (
	"fmt"
	"strconv"

	"github.com/NeGat1FF/task-cli/task"
	"github.com/spf13/cobra"
)

// markCmd represents the mark command
var markCmd = &cobra.Command{
	Use:   "mark",
	Short: "Mark a task with a specific tag.",
	Long:  `Mark a task with the specified tag. This command allows you to update the status of a task by specifying its ID and the tag you want to apply. Tags correspond to task statuses such as 'TODO', 'In Progress', and 'Done'. This is a base command for marking tasks, and more specific subcommands are available for each status.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

var markDoneCmd = &cobra.Command{
	Use:   "done [task ID]",
	Short: "Mark a task as Done.",
	Long:  "Mark the task with the specified ID as 'Done'. Provide the task ID as a positional argument. This command updates the status of the task to 'Done', indicating that it has been completed. If the task with the given ID is found and updated successfully, a confirmation message will be displayed.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		val, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}

		err = task.MarkTask(val, task.Done)
		if err != nil {
			return err
		}

		fmt.Printf("Task with ID: %d marked as Done", val)

		return nil
	},
}

var markInProgressCmd = &cobra.Command{
	Use:   "in-progress [task ID]",
	Short: "Mark a task as In Progress.",
	Long:  "Mark the task with the specified ID as 'In Progress'. Provide the task ID as a positional argument. This command updates the status of the task to 'In Progress', indicating that work on the task is currently underway. If the task with the given ID is found and updated successfully, a confirmation message will be displayed.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		val, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}

		err = task.MarkTask(val, task.InProgress)
		if err != nil {
			return err
		}

		fmt.Printf("Task with ID: %d marked as in-progress", val)

		return nil
	},
}

func init() {
	markCmd.AddCommand(markDoneCmd)
	markCmd.AddCommand(markInProgressCmd)

	rootCmd.AddCommand(markCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// markCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// markCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
