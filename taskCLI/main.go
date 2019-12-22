package main

import (
	"github.com/amitlevy21/gophercises/taskCLI/cmd"
	"github.com/amitlevy21/gophercises/taskCLI/db"
)

func main() {
	db.Init()
	cmd.Execute()
}
