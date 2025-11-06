package cmd

import (
	"github.com/henryyr/gitprofile/internal/profile"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new profile",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := profile.AddProfile(); err != nil {
			return err
		}
		return nil
	},
}
