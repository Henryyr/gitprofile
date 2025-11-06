package cmd

import (
	"fmt"

	"github.com/henryyr/gitprofile/internal/profile"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all profiles",
	RunE: func(cmd *cobra.Command, args []string) error {
		profiles, err := profile.LoadProfiles()
		if err != nil {
			return err
		}

		if len(profiles.Entries) == 0 {
			fmt.Println("No profiles found.")
			return nil
		}

		for _, p := range profiles.Entries {
			if p.Username == profiles.Active {
				fmt.Printf("✅ %s <%s> (%s)\n", p.Username, p.Email, p.Name)
			} else {
				fmt.Printf("  %s <%s> (%s)\n", p.Username, p.Email, p.Name)
			}
		}

		return nil
	},
}
