package config

import (
	"path/filepath"

	"runtime"

	"github.com/hadenlabs/docker-kali-linux/internal/errors"
	"github.com/hadenlabs/docker-kali-linux/internal/version"
)

const tagLatest string = "latest"

// Config struct field.
type Config struct {
	Docker Docker
	App    App
}

// ReadConfig read values and files for config.
func (c *Config) ReadConfig() (*Config, error) {
	_, filePath, _, _ := runtime.Caller(0)
	rootDir := filepath.Dir(filePath)
	c.App.RootPath = filepath.Join(rootDir, "..")
	tag := version.Short()
	c.Docker.TargetImage = targetImage
	c.Docker.Tag = tag
	c.Docker.TagLatest = tagLatest
	c.App.Version = tag
	return c, nil
}

// Initialize new instance.
func Initialize() *Config {
	conf := New()
	conf, err := conf.ReadConfig()
	if err != nil {
		panic(errors.Wrap(err, errors.ErrorReadConfig, ""))
	}
	return conf
}

// New create config.
func New() *Config {
	return &Config{}
}
