package profile

import (
	"encoding/json"
	"errors"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/henryyr/gitprofile/util"
)

const configFileName = "profiles.json"

func getConfigFilePath() (string, error) {
	dir, err := util.GetConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, configFileName), nil
}

func LoadProfiles() (*Profiles, error) {
	path, err := getConfigFilePath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return &Profiles{Entries: []Profile{}}, nil
	}
	if err != nil {
		return nil, err
	}

	var profiles Profiles
	if err := json.Unmarshal(data, &profiles); err != nil {
		return nil, err
	}
	return &profiles, nil
}

func SaveProfiles(p *Profiles) error {
	path, err := getConfigFilePath()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

func SetGitConfig(username, email string) error {
	cmds := [][]string{
		{"git", "config", "--global", "user.name", username},
		{"git", "config", "--global", "user.email", email},
	}

	for _, cmd := range cmds {
		if err := exec.Command(cmd[0], cmd[1:]...).Run(); err != nil {
			return err
		}
	}
	return nil
}
