package cmd

import (
	"fmt"

	"github.com/henryyr/gitprofile/internal/profile"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(currentCmd)
}

var currentCmd = &cobra.Command{
	Use:   "current",
	Short: "Show the current active profile",
	RunE: func(cmd *cobra.Command, args []string) error {
		profiles, err := profile.LoadProfiles()
		if err != nil {
			return err
		}

		if profiles.Active == "" {
			fmt.Println("No active profile.")
			return nil
		}

		for _, p := range profiles.Entries {
			if p.Username == profiles.Active {
				fmt.Printf("Current profile: %s <%s> (%s)\n", p.Username, p.Email, p.Name)
				return nil
			}
		}

		return fmt.Errorf("active profile '%s' not found", profiles.Active)
	},
}
