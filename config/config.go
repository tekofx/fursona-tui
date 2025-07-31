package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
)

type Config struct {
	Name      string            `json:"name"`
	Surname   string            `json:"surname"`
	Species   string            `json:"species"`
	Gender    string            `json:"gender"`
	Pronouns  string            `json:"pronouns"`
	Palette   []string          `json:"palette"`
	OtherData map[string]string `json:"otherData"`
}

func GetConfigPath() string {
	var configDir string

	if runtime.GOOS == "windows" {
		dir, err := os.UserConfigDir()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to get config directory:", err)
			dir, err = os.UserHomeDir()
			if err != nil {
				fmt.Fprintln(os.Stderr, "Failed to get home directory:", err)
				return ""
			}
			return filepath.Join(dir, "fursona-tui")
		}
		configDir = filepath.Join(dir, "fursona-tui")
	} else {
		dir, err := os.UserHomeDir()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to get home directory:", err)
			return ""
		}
		configDir = filepath.Join(dir, ".config", "fursona-tui")
	}

	return configDir
}

func DefaultConfig() Config {

	palette := [2]string{"#FFFFFF", "#000000"}
	return Config{
		Name:     "Name",
		Surname:  "Surname",
		Species:  "Species",
		Gender:   "Gender",
		Pronouns: "Pronouns",
		Palette:  palette[:],
		OtherData: map[string]string{
			"favorite_food": "pizza",
			"hobby":         "drawing",
		},
	}
}

func ReadConfig() *Config {
	configPath := GetConfigPath()
	settingsFilePath := filepath.Join(configPath, "settings.json")

	// Create config dir if it does not exist
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		if err := os.MkdirAll(configPath, 0755); err != nil {
			fmt.Fprintln(os.Stderr, "Failed to create config directory:", err)
			config := DefaultConfig()
			return &config
		}
	}

	// Check if config file exists
	if _, err := os.Stat(settingsFilePath); os.IsNotExist(err) {
		// Create default config
		defaultConfig := DefaultConfig()
		file, err := os.Create(settingsFilePath)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to create config file:", err)
			return &defaultConfig
		}
		defer file.Close()
		encoder := json.NewEncoder(file)
		encoder.SetIndent("", "  ")
		if err := encoder.Encode(defaultConfig); err != nil {
			fmt.Fprintln(os.Stderr, "Failed to write default config:", err)
			return &defaultConfig
		}

		fmt.Printf("No config detected, config created at %s.\n", settingsFilePath)
		fmt.Println("Please edit the configuration file to modify the data showed.")
		os.Exit(0)
	}

	file, err := os.Open(settingsFilePath)
	if err != nil {
		return nil
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil
	}

	var config Config
	if err := json.Unmarshal(bytes, &config); err != nil {
		return nil
	}
	return &config
}
