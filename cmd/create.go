/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/0xk2/todocli/pkg"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a task",
	Long:  `just create a task`,
	Run: func(cmd *cobra.Command, args []string) {
		taskTitle, _ := cmd.Flags().GetString("title")
		taskDescription, _ := cmd.Flags().GetString("description")
		t := pkg.NewTask(taskTitle, taskDescription)
		fmt.Println("Task ID: ", t.ID())
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("title", "t", "", "specify task title")
	createCmd.Flags().StringP("description", "d", "", "specify task description")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
