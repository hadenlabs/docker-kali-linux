package test

import (
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/docker"
	"github.com/stretchr/testify/assert"

	"fmt"

	"github.com/hadenlabs/docker-kali-linux/config"
)

func TestToolsValidateScanSuccess(t *testing.T) {
	conf := config.Initialize()
	imageTag := fmt.Sprintf("%s:%s", conf.Docker.TargetImage, "base")
	otherOptions := []string{}
	expectApps := []string{
		"amap",
		"commix",
		"dirb",
		"dnschef",
		"dnsenum",
		"dnsmap",
		"dnsrecon",
		"dnswalk",
		"hping3",
		"hydra",
		"msfconsole",
		"msfvenom",
		"ncrack",
		"netdiscover",
		"nikto",
		"nmap",
		"padbuster",
		"reaver",
		"responder",
		"searchsploit",
		"socat",
		"telnet",
		"tshark",
		"wash",
		"wireshark",
		"zaproxy",
	}

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
	assert.Subset(t, strings.Split(outputListApps, "\n"), expectApps)
}
