package cmd

import (
	"fmt"

	"github.com/henryyr/gitprofile/internal/profile"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(renameCmd)
}

var renameCmd = &cobra.Command{
	Use:   "rename [oldUsername] [newUsername]",
	Short: "Rename an existing profile",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		oldUsername := args[0]
		newUsername := args[1]

		profiles, err := profile.LoadProfiles()
		if err != nil {
			return err
		}

		// Cek apakah newUsername udah ada
        for _, p := range profiles.Entries {
            if p.Username == newUsername {
                return fmt.Errorf("❌ profile '%s' already exists", newUsername)
            }
        }

		var found bool
		for i, p := range profiles.Entries {
			if p.Username == oldUsername {
				profiles.Entries[i].Username = newUsername
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("❌ profile '%s' not found", oldUsername)
		}

		// update active kalau perlu
		if profiles.Active == oldUsername {
			profiles.Active = newUsername
		}

		if err := profile.SaveProfiles(profiles); err != nil {
			return err
		}

		fmt.Printf("✅ Renamed profile '%s' → '%s'\n", oldUsername, newUsername)
		return nil
	},
}
