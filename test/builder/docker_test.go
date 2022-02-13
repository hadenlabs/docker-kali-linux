package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/docker"
	"github.com/stretchr/testify/assert"

	"fmt"

	"github.com/hadenlabs/docker-kali-linux/config"
)

func TestToolsBuilderSuccess(t *testing.T) {
	conf := config.Initialize()
	imageTag := fmt.Sprintf("%s:%s", conf.Docker.TargetImage, "latest")
	otherOptions := []string{}
	buildOptions := &docker.BuildOptions{
		Tags:         []string{imageTag},
		OtherOptions: otherOptions,
	}

	docker.Build(t, conf.App.RootPath, buildOptions)
	opts := &docker.RunOptions{
		Command: []string{
			"bash", "-c",
			"compgen -c", "|",
			"sort -u",
		},
	}
	outputListApps := docker.Run(t, imageTag, opts)
	assert.NotEmpty(t, outputListApps, outputListApps)
}

func TestGetDockerBuilderSuccess(t *testing.T) {
	conf := config.Initialize()
	imageTag := fmt.Sprintf("%s:%s", conf.Docker.TargetImage, "latest")
	otherOptions := []string{}
	buildOptions := &docker.BuildOptions{
		Tags:         []string{imageTag},
		OtherOptions: otherOptions,
	}

	docker.Build(t, conf.App.RootPath, buildOptions)
	opts := &docker.RunOptions{
		Command: []string{
			"bash", "-c",
			"compgen -c", "|",
			"sort -u",
		},
	}
	outputListApps := docker.Run(t, imageTag, opts)
	assert.NotEmpty(t, outputListApps, outputListApps)
}
