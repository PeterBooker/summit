package config

import (
	"log"
	"os"
)

// Config contains global application information
type Config struct {
	Version string
	Name    string
	Commit  string
	Date    string
	WD      string
	Port    string
}

// Setup creates, fills and returns the Config struct
func Setup(version, commit, date string) *Config {

	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting Working Directory: %s\n", err)
	}

	config := &Config{
		Version: version,
		Name:    "Summit",
		Commit:  commit,
		Date:    date,
		WD:      wd,
		Port:    ":4000",
	}

	return config
}
