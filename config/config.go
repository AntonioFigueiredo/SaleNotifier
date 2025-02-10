package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func findProjectRoot(startDir string) (string, error) {
	for {
		if _, err := os.Stat(filepath.Join(startDir, ".git")); err == nil {
			return startDir, nil
		}

		parentDir := filepath.Dir(startDir)
		if parentDir == startDir {
			break // We've reached the root of the filesystem without finding .git
		}
		startDir = parentDir
	}

	return "", fmt.Errorf(".git directory not found")
}

func Init() {
	// Get the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		log.Error().Err(err).Msg("Error getting current directory")
		return
	}

	// Find the project root directory (the directory containing .git)
	projectRoot, err := findProjectRoot(currentDir)
	if err != nil {
		log.Error().Err(err).Msg("Could not find project root directory")
		return
	}

	// Construct the absolute path to the .env file in the project root
	envFilePath := filepath.Join(projectRoot, ".env")

	// Load the .env file
	if err := godotenv.Load(envFilePath); err != nil {
		log.Error().Err(err).Msg("Could not load .env file")
	}
}
