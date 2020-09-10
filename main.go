package main

import (
	"fmt"
	"os"
	"uploader/internal/cli"
	"uploader/internal/db"
)

func main() {
	must(cli.RootCmd.Execute())
	defer db.DB.Close()

}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
