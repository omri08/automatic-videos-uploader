package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"uploader/services"
)

// runCmd represents the upload command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run tell the program to start upload the videos",
	Run: func(cmd *cobra.Command, args []string) {
		lessons := services.LoadLessons()
		if lessons == nil {
			fmt.Printf("Error loading Lessons\n")

		} else {
			list := services.ListVideos(`C:\Users\Omri\Desktop\records`, lessons)
			services.UploadToYoutube(list)
		}
	},
}

func init() {
	RootCmd.AddCommand(runCmd)
}
