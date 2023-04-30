package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/xV0lk/cli-test/cmd"
	"github.com/xV0lk/cli-test/db"
)

func main() {
	// get the homedir on any os
	hd, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	// get db file path
	dbPath := filepath.Join(hd, "tasks.db")
	// use the path to connect to db and set it to the global db var
	// It creates a db if it does't exist
	must(db.InitDb(dbPath))
	fmt.Printf("Connected to db tasks on %s.\n", dbPath)
	must(cmd.RootCmd.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
