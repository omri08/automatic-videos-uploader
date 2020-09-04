package cmd

import (
	"github.com/spf13/cobra"
	"uploader/services"
)

// runCmd represents the upload command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run tell the program to start upload the videos",
	Run: func(cmd *cobra.Command, args []string) {
		lessons := services.LoadLessons()
		list := services.ListVideos(`C:\Users\Omri\Desktop\records`, lessons)
		services.UploadToYoutube(list)
	},
}

func init() {
	RootCmd.AddCommand(runCmd)
}
