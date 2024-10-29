/*
Copyright © 2024 Ethan Grams ethandoescode@gmail.com 
*/
package cmd

import (
	"fmt"
	"log"
	"os"
    "encoding/csv"
    "text/tabwriter"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "View todo list",
    Long: "",
	Run: func(cmd *cobra.Command, args []string) {
        err := displayTodoList("todo.csv")
        if err != nil {
            log.Fatal("No task list. Create some tasks with the add command!")
        }
    },
}

func displayTodoList(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()
    
    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        return err
    }
    
    w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.AlignRight)
    fmt.Fprintln(w, "   Task\t      Time\t")
    fmt.Fprintln(w, "---------------------")

    for _, record := range records {
        fmt.Fprintf(w, "%s\t%s\t\n", record[0], record[1])
    }
    w.Flush()
    
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
