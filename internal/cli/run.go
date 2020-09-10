package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"uploader/internal/db"
	"uploader/pkg"
)

// runCmd represents the upload command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run tell the program to start upload the videos",
	Run: func(cmd *cobra.Command, args []string) {
		lessons := db.DB.LoadLessons()
		if lessons == nil {
			fmt.Printf("Error loading Lessons\n")

		} else {
			list := pkg.ListVideos(`C:\Users\Omri\Desktop\records`, lessons)
			pkg.UploadToYoutube(list)
		}
	},
}

func init() {
	RootCmd.AddCommand(runCmd)
}
