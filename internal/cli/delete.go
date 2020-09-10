package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
	"uploader/internal/db"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a lesson from the json file",

	Run: func(cmd *cobra.Command, args []string) {
		lesson := strings.Join(args, " ")
		err := db.DB.DeleteLesson(lesson)
		if err != nil {
			fmt.Printf("%s deleted successfully\n", lesson)
		} else {
			fmt.Printf("failed to deleted %s\n", lesson)
		}
	},
}

func init() {
	RootCmd.AddCommand(deleteCmd)
}
