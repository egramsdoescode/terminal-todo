/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove [task]",
	Short: "Remove a task",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
	    if len(args) == 1{
            task := args[0]

            err := removeTask(task)
            if err != nil {
                log.Fatal(err)
            }
        } else {
            log.Println("Invalid number of arguments")
        }
    },
}

func removeTask(task string) error {
    linesToKeep, err := makeNewList(task)
    if err != nil {
        return err
    }

    file, err := os.OpenFile("todo.csv", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
    if err != nil {
        return err
    }
    defer file.Close()
    
    for _, line := range linesToKeep {
        _, err := file.WriteString(line + "\n")
        if err != nil {
            return err
        }
    }

    return nil
}

func makeNewList(task string) ([]string, error) {
	file, err := os.Open("todo.csv") 
	if err != nil {
		return nil, err
	}

	defer file.Close()
    
    var linesToKeep []string

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        data := strings.Split(scanner.Text(), ",")
        if data[0] != task {
            linesToKeep = append(linesToKeep, scanner.Text())
        }
    }

    err = scanner.Err()
    if err != nil {
        return nil, err
    }

    return linesToKeep, nil 
}

func init() {
	rootCmd.AddCommand(removeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
