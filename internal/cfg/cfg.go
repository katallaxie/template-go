package cfg

import (
	"os"

	"github.com/katallaxie/go-template/pkg/spec"
)

// Config is the configuration for the application.
type Config struct {
	// Verbose enables verbose logging.
	Verbose bool
	// Force forces the creation of the file.
	Force bool
	// File is the file to write to.
	File string
	// Root is the root configuration.
	Root Root
}

// Root is the root configuration.
type Root struct {
	// Hook is the name of the hook to run.
	Run string
}

// Cwd is the current working directory.
func (c *Config) Cwd() (string, error) {
	p, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return p, nil
}

// New returns a new Config.
func New() *Config {
	return &Config{}
}

// Default returns the default configuration.
func Default() *Config {
	return &Config{
		File: spec.DefaultFileName,
	}
}
