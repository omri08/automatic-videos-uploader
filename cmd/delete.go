package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
	"uploader/services"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a lesson from the json file",

	Run: func(cmd *cobra.Command, args []string) {
		lesson := strings.Join(args, " ")
		deleted := services.DeleteLesson(lesson)
		if deleted {
			fmt.Printf("%s deleted successfully\n", lesson)
		} else {
			fmt.Printf("failed to deleted %s\n", lesson)
		}
	},
}

func init() {
	RootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
