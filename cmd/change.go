/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/0xk2/todocli/pkg"
	"github.com/spf13/cobra"
)

// changeCmd represents the change command
var changeCmd = &cobra.Command{
	Use:   "change",
	Short: "Command to change state of a task",
	Long: `Code is:
1 to start / doing
0 to pause
2 to finish
3 to stop
Note: You can not change a stopped or finished Task
`,
	Run: func(cmd *cobra.Command, args []string) {
		taskId, _ := cmd.Flags().GetString("id")
		code, err := cmd.Flags().GetInt8("status")
		if err != nil {
			fmt.Println("cannot get code")
		}
		t := pkg.LoadTask(taskId)
		t.ChangeStatus(code)
	},
}

func init() {
	rootCmd.AddCommand(changeCmd)

	changeCmd.Flags().StringP("id", "i", "", "Task ID")
	changeCmd.Flags().Int8P("status", "s", 0, "Code to dictate the next state")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// changeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// changeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
