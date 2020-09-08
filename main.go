package main

import (
	"fmt"
	"os"
	"uploader/cmd"
	"uploader/db"
)

func main() {

	must(cmd.RootCmd.Execute())
	defer db.DB.Close()

}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
