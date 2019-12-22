package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/amitlevy21/gophercises/taskCLI/db"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
		task := strings.Join(args[:], " ")
		if err := db.Add(task); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
