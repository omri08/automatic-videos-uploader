package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"uploader/db"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all your lessons",
	Run: func(cmd *cobra.Command, args []string) {
		lessonArr := db.GlobalMySQL.LoadLessons()
		for i, lesson := range lessonArr {
			fmt.Printf("%d) name: %s, day: %d, starts: %d, ends: %d\n", i+1, lesson.Name, lesson.Day+1, lesson.Starts, lesson.Ends)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)

}
