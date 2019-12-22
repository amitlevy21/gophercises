package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/amitlevy21/gophercises/taskCLI/db"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do [task_id]",
	Short: "Mark task as completed",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
		if err := doTask(args); err != nil {
			fmt.Println(err)
		}
	},
}

func doTask(ids []string) error {
	for _, id := range ids {
		fmt.Println(id)
		taskId, err := strconv.Atoi(id)
		if err != nil {
			return err
		}
		err = db.Delete(taskId)
		if err != nil {
			return err
		}
	}
	return nil
}
