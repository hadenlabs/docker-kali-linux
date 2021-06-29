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
	imageTag := fmt.Sprintf("%s:%s", conf.Docker.TargetImage, "base")
	otherOptions := []string{}
	buildOptions := &docker.BuildOptions{
		Tags:         []string{imageTag},
		OtherOptions: otherOptions,
	}

	docker.Build(t, "../../images/base", buildOptions)
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
	imageTag := fmt.Sprintf("%s:%s", conf.Docker.TargetImage, "base")
	otherOptions := []string{}
	buildOptions := &docker.BuildOptions{
		Tags:         []string{imageTag},
		OtherOptions: otherOptions,
	}

	docker.Build(t, "../../images/base", buildOptions)
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
