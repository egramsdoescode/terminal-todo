/*
Copyright © 2024 Ethan Grams ethandoescode@gmail.com 
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "View todo list",
    Long: "",
	Run: func(cmd *cobra.Command, args []string) {
        err := readTodoList()
        if err != nil {
            log.Fatal("No task list. Create some tasks with the add command!")
        }
    },
}

func readTodoList() error {
    list, err := os.ReadFile("todo.csv")
    if err != nil {
        return err
    }
    
    if len(list) != 0 { 
        fmt.Println(string(list))
    } else {
        fmt.Println("No tasks")
    }
    return nil
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
