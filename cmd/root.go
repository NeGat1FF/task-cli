package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "task-cli",
	Short: "Manage your tasks with this CLI tool.",
	Long: `This CLI tool helps you manage your tasks efficiently. With commands to add, delete, list, and update tasks, you can keep track of your tasks and their statuses. Whether you need to mark tasks as 'TODO', 'In Progress', or 'Done', this tool provides an easy-to-use interface for task management.

Examples of usage include:

task-cli add [task description] – Add a new task with the specified description.
task-cli list – List all tasks.
task-cli delete [task ID] – Delete a task by its ID.
task-cli mark done [task ID] – Mark a task as 'Done'.
This application uses the Cobra library for creating command-line interfaces in Go, providing a structured and extensible way to interact with your task list.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.task-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
