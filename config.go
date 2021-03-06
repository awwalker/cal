package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
)

// Details stored inside the conig file.
type config struct {
	Version   string              `json:"version"`
	Calendars map[string]calendar `json:"calendars"`
}

// newConfig - create a new config.
func newConfig() *config {
	cfg := new(config)
	cfg.Version = "0.0.1"
	// Empty Calendars
	cfg.Calendars = make(map[string]calendar)
	return cfg
}

// loadCalConfig - if there is a prexisting config load it.
func loadCalConfig() (*config, error) {
	// Check config path.
	configPath, err := getCalConfigPath()
	if err != nil {
		return nil, err
	}
	// Read the config.
	configBytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	var cfg *config
	if err := json.Unmarshal(configBytes, &cfg); err != nil {
		return nil, err
	}
	// Success.
	return cfg, nil
}

// saveCalConfig - save an existing config.
func saveCalConfig(config *config) error {
	// Create a new configDir if it does not already exist.
	if err := createCalConfigDir(); err != nil {
		return err
	}
	path, err := getCalConfigPath()
	if err != nil {
		return err
	}
	jsonBytes, err := json.Marshal(&config)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(path, jsonBytes, 0700); err != nil {
		return err
	}
	return nil
}

// getCalConfigDir - get default Config directory location.
// TODO: (@hc1334) allow user defined Config directory location.
func getCalConfigDir() (string, error) {
	curUser, err := user.Current()
	if err != nil {
		return "", err
	}
	homeDir := curUser.HomeDir
	if homeDir == "" {
		return "", fmt.Errorf("No Home dir found for current user")
	}

	// Create Cal Config dir Path.
	// TODO: (@hc1334) differentiate windows OS.
	configDir := filepath.Join(homeDir, globalCalConfigDir)
	return configDir, nil
}

// createCalAliasDir - create a dir for each calendar aliased in the config.
func createCalAliasDir(alias string) error {
	configDir, err := getCalConfigDir()
	if err != nil {
		return err
	}
	// Create new dir under config dir for each calendar alias.
	path := filepath.Join(configDir, alias)
	return os.MkdirAll(path, 0700)
}

// getCalAliasDir - given a calendar alias get the directory it's client_id.json is stored in.
func getCalAliasDir(alias string) (string, error) {
	calDir, err := getCalConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(calDir, alias), nil
}

// Create a new Config directory
func createCalConfigDir() error {
	configDir, err := getCalConfigDir()
	if err != nil {
		return err
	}
	// Create new Cal Config Dir.
	if err = os.MkdirAll(configDir, 0700); err != nil {
		return err
	}
	return nil
}

// Get the path to Cal Config.
func getCalConfigPath() (string, error) {
	dir, err := getCalConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, globalCalConfigFile), nil
}

// isCalConfigExists - Verifies is a config currently exists.
func isCalConfigExists() bool {
	configPath, err := getCalConfigPath()
	if err != nil {
		return false
	}
	if _, err = os.Stat(configPath); err != nil {
		return false
	}
	return true
}

// TODO: (@hc1334) setCalendar method belonging to *config.
func (c *config) addCalendar(cal calendar) {
	if _, ok := c.Calendars[cal.Alias]; !ok {
		c.Calendars[cal.Alias] = cal
	}
}
