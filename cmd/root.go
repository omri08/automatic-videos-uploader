package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "uploader",
	Short: "Uploader is a CLI tool to automatically upload videos to youtube ",
}
