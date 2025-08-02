package nerdctl

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

type nerdctlCli struct {
	filePath string
}

// ContainerInspect
func (p *nerdctlCli) ContainerInspect(name string) (string, error) {
	return p.exec("inspect", name)
}

// ContainerList
func (p *nerdctlCli) ContainerList() (string, error) {
	return p.exec("container", "list", "-a")
}

// ContainerLogs
func (p *nerdctlCli) ContainerLogs(name string) (string, error) {
	return p.exec("logs", name)
}

// ContainerRemove
func (p *nerdctlCli) ContainerRemove(name string) (string, error) {
	return p.exec("container", "rm", name)
}

// ContainerRun
func (p *nerdctlCli) ContainerRun(imageName string, portMappings map[int]int, envVariables []string) (string, error) {
	args := []string{"run", "--rm", "-d"}
	if len(portMappings) > 0 {
		for hostPort, containerPort := range portMappings {
			args = append(args, fmt.Sprintf("--publish=%d:%d", hostPort, containerPort))
		}
	} else {
		args = append(args, "--publish-all")
	}
	for _, env := range envVariables {
		args = append(args, "--env", env)
	}
	output, err := p.exec(append(args, imageName)...)
	if err == nil {
		return output, nil
	}
	if strings.Contains(output, "Error: short-name") {
		imageName = "docker.io/" + imageName
		if output, err = p.exec(append(args, imageName)...); err == nil {
			return output, nil
		}
	}
	return "", err
}

// ContainerStop
func (p *nerdctlCli) ContainerStop(name string) (string, error) {
	return p.exec("container", "stop", name)
}

// ImageBuild
func (p *nerdctlCli) ImageBuild(containerFile string, imageName string) (string, error) {
	args := []string{"build"}
	if imageName != "" {
		args = append(args, "-t", imageName)
	}
	return p.exec(append(args, "-f", containerFile)...)
}

// ImageList
func (p *nerdctlCli) ImageList() (string, error) {
	return p.exec("images", "--digests")
}

// ImagePull
func (p *nerdctlCli) ImagePull(imageName string) (string, error) {
	output, err := p.exec("image", "pull", imageName)
	if err == nil {
		return fmt.Sprintf("%s\n%s pulled successfully", output, imageName), nil
	}
	if strings.Contains(output, "Error: short-name") {
		imageName = "docker.io/" + imageName
		if output, err = p.exec("pull", imageName); err == nil {
			return fmt.Sprintf("%s\n%s pulled successfully", output, imageName), nil
		}
	}
	return "", err
}

// ImagePush
func (p *nerdctlCli) ImagePush(imageName string) (string, error) {
	output, err := p.exec("image", "push", imageName)
	if err == nil {
		return fmt.Sprintf("%s\n%s pushed successfully", output, imageName), nil
	}
	return "", err
}

// ImageRemove
func (p *nerdctlCli) ImageRemove(imageName string) (string, error) {
	return p.exec("image", "rm", imageName)
}

// NetworkList
func (p *nerdctlCli) NetworkList() (string, error) {
	return p.exec("network", "ls")
}

// VolumeList
func (p *nerdctlCli) VolumeList() (string, error) {
	return p.exec("volume", "ls")
}

func (p *nerdctlCli) exec(args ...string) (string, error) {
	output, err := exec.Command(p.filePath, args...).CombinedOutput()
	return string(output), err
}

func newNerdctlCli() (*nerdctlCli, error) {
	for _, cmd := range []string{"nerdctl", "nerdctl.exe"} {
		filePath, err := exec.LookPath(cmd)
		if err != nil {
			continue
		}
		if _, err = exec.Command(filePath, "version").CombinedOutput(); err == nil {
			return &nerdctlCli{filePath}, nil
		}
	}
	return nil, errors.New("nerdctl CLI not found")
}
