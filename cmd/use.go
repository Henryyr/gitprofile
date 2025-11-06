package cmd

import (
	"fmt"

	"github.com/henryyr/gitprofile/internal/profile"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(useCmd)
}

var useCmd = &cobra.Command{
	Use:   "use [username]",
	Short: "Switch to a different profile",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		username := args[0]

		profiles, err := profile.LoadProfiles()
		if err != nil {
			return err
		}

		for _, p := range profiles.Entries {
			if p.Username == username {
				if err := profile.SetGitConfig(p.Username, p.Email); err != nil {
					return err
				}
				profiles.Active = username
				if err := profile.SaveProfiles(profiles); err != nil {
					return err
				}
				fmt.Printf("✅ Switched to profile: %s\n", username)
				return nil
			}
		}

		return fmt.Errorf("profile '%s' not found", username)
	},
}
