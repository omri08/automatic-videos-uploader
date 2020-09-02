package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"uploader/models"
	"uploader/services"
)

var (
	name              string
	day, starts, ends int

	addCmd = &cobra.Command{
		Use:   "add",
		Short: "Adds a lesson to your schedule  ",

		Run: func(cmd *cobra.Command, args []string) {
			l := models.Lesson{Name: name, Day: day - 1, Starts: starts, Ends: ends}

			if err := services.AddLessonToJSON(l); err != nil {
				fmt.Printf("%v\n", err)
			} else {
				fmt.Printf("Lesson added successfully")
			}

		},
	}
)

func init() {
	addCmd.Flags().StringVar(&name, "name", "foo", "The name of the lesson")
	addCmd.Flags().IntVar(&day, "day", 1, "The numric day of the lesson")
	addCmd.Flags().IntVar(&starts, "starts", 13, "The hour the lesson starts")
	addCmd.Flags().IntVar(&ends, "ends", 17, "The hour the lesson ends")
	RootCmd.AddCommand(addCmd)
}
