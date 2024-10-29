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

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [task] [time]",
	Short: "Add a task",
    Long: "A task that does some things",
	Run: func(cmd *cobra.Command, args []string) {
        if len(args) == 2 {
            task := args[0]
            time := args[1]
        
            err := addTask(task, time)
            if err != nil {
                fmt.Printf("error: %v\n", err)
            } else {
                fmt.Println("Task added successfully")
            }
        } else {
            log.Fatal("Invalid arguments")
        }
    },
}

func addTask(task string, time string) error {
    formattedTask := fmt.Sprintf("%s,%s\n", task, time)

    file, err := os.OpenFile("todo.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
    if err != nil {
        return err 
    }
    defer file.Close()
    
    item := []byte(formattedTask)

    _, err = file.Write(item)
    if err != nil {
        return err
    }

    return nil 
}

func init() {
	rootCmd.AddCommand(addCmd)
}
