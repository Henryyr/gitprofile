package cmd

import (
	"fmt"

	"github.com/henryyr/gitprofile/internal/profile"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(removeCmd)
}

var removeCmd = &cobra.Command{
	Use:   "remove [username]",
	Short: "Remove a profile",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		username := args[0]

		profiles, err := profile.LoadProfiles()
		if err != nil {
			return err
		}

		var newEntries []profile.Profile
		var found bool
		for _, p := range profiles.Entries {
			if p.Username == username {
				found = true
			} else {
				newEntries = append(newEntries, p)
			}
		}

		if !found {
			return fmt.Errorf("profile '%s' not found", username)
		}

		profiles.Entries = newEntries

		if profiles.Active == username {
			profiles.Active = ""
		}

		if err := profile.SaveProfiles(profiles); err != nil {
			return err
		}

		fmt.Printf("✅ Profile removed: %s\n", username)
		return nil
	},
}
