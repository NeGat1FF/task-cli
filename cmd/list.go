package cmd

import (
	"fmt"

	"github.com/NeGat1FF/task-cli/task"
	"github.com/spf13/cobra"
)

func PrintTasks(n int) error {
	tasks, err := task.GetTasks(n)
	if err != nil {
		return err
	}
	for _, task := range tasks {
		fmt.Println(task)
	}

	return nil
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long:  `List all tasks you added to the task tracker. This command displays every task in the task list. You can use flags to filter tasks by their status, for example, task-cli list todo will list only tasks from the 'TODO' category`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return PrintTasks(0)
	},
}

var listTodoCmd = &cobra.Command{
	Use:   "todo",
	Short: "List tasks marked as TODO",
	Long:  "List tasks that are marked with the status 'TODO'. This command filters and displays only the tasks that are currently categorized as 'TODO'",
	RunE: func(cmd *cobra.Command, args []string) error {
		return PrintTasks(task.Todo)
	},
}

var listDoneCmd = &cobra.Command{
	Use:   "done",
	Short: "List tasks marked as Done",
	Long:  "List tasks that are marked with the status 'Done'. This command filters and displays only the tasks that have been completed and are categorized as 'Done'",
	RunE: func(cmd *cobra.Command, args []string) error {
		return PrintTasks(task.Done)
	},
}

var listInProgressCmd = &cobra.Command{
	Use:   "in-progress",
	Short: "List tasks that are marked with the status 'In Progress'. This command filters and displays only the tasks that are currently ongoing and categorized as 'In Progress'",
	RunE: func(cmd *cobra.Command, args []string) error {
		return PrintTasks(task.InProgress)
	},
}

func init() {
	listCmd.AddCommand(listTodoCmd)
	listCmd.AddCommand(listDoneCmd)
	listCmd.AddCommand(listInProgressCmd)

	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
