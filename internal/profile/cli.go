package profile

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AddProfile() error {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("username: ")
	username, _ := reader.ReadString('\n')

	fmt.Print("email: ")
	email, _ := reader.ReadString('\n')

	fmt.Print("name: ")
	name, _ := reader.ReadString('\n')

	username = strings.TrimSpace(username)
	email = strings.TrimSpace(email)
	name = strings.TrimSpace(name)

	profiles, err := LoadProfiles()
	if err != nil {
		return err
	}

	// cek duplikat
	for _, p := range profiles.Entries {
		if p.Username == username {
			return fmt.Errorf("profile '%s' already exists", username)
		}
	}

	newProfile := Profile{
		Name:     name,
		Email:    email,
		Username: username,
	}

	profiles.Entries = append(profiles.Entries, newProfile)

	if err := SaveProfiles(profiles); err != nil {
		return err
	}

	fmt.Println("✅ Profile added:", username)
	return nil
}
