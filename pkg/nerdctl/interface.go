package nerdctl

// Nerdctl interface
type Nerdctl interface {
	// ContainerInspect displays the low-level information on containers identified by the ID or name
	ContainerInspect(name string) (string, error)
	// ContainerList lists all the containers on the system
	ContainerList() (string, error)
	// ContainerLogs Display the logs of a container
	ContainerLogs(name string) (string, error)
	// ContainerRemove removes a container
	ContainerRemove(name string) (string, error)
	// ContainerRun pulls an image from a registry
	ContainerRun(imageName string, portMappings map[int]int, envVariables []string) (string, error)
	// ContainerStop stops a running container using the ID or name
	ContainerStop(name string) (string, error)
	// ImageBuild builds an image from a Dockerfile
	ImageBuild(dockerFile string, imageName string) (string, error)
	// ImageList list the container images on the system
	ImageList() (string, error)
	// ImagePull pulls an image from a registry
	ImagePull(imageName string) (string, error)
	// ImagePush pushes an image to a registry
	ImagePush(imageName string) (string, error)
	// ImageRemove removes an image from the system
	ImageRemove(imageName string) (string, error)
	// NetworkList lists all the networks on the system
	NetworkList() (string, error)
	// VolumeList lists all the volumes on the system
	VolumeList() (string, error)
}

func NewNerdctl() (Nerdctl, error) {
	// TODO: add implementations for Nerdctl bindings and Docker CLI
	return newNerdctlCli()
}
