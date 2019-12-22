package cmd

import (
	"fmt"
	"github.com/amitlevy21/gophercises/taskCLI/db"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List unfinished tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks := db.Tasks()
		for i, t := range tasks {
			fmt.Printf("%d. %s\n", i+1, t.Name)
		}
	},
}
