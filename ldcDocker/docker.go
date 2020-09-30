package ldcDocker

import (
	"github.com/jeyvison/liferay-docker-control/ldcLog"
	"os/exec"
)

type DockerControl struct {
}

func getDockerExecutablePath() string {
	return "/usr/local/bin/docker"
}

var logManager = ldcLog.DefaultLogManager

var logger = logManager.Logger

var loggerWriter = logManager.LoggerWriter

func (dockerControl DockerControl) StopContainer(containerName string) {
	dockerExecutablePath := getDockerExecutablePath()

	dockerStopContainerCommand := []string{dockerExecutablePath, "stop", containerName}

	logger.Println("Stopping container " + containerName)

	dockerCommand := &exec.Cmd{
		Path:   dockerExecutablePath,
		Args:   dockerStopContainerCommand,
		Stdout: loggerWriter,
		Stderr: loggerWriter,
	}

	dockerCommand.Run()
}

func (dockerControl DockerControl) RunDocker(containerName string, imageName string, portMapping string) error {

	dockerExecutablePath := getDockerExecutablePath()

	logger.Println("Stopping container " + containerName)

	dockerControl.StopContainer(containerName)

	logger.Println("Removing container " + containerName)

	dockerRemoveContainerCommand := []string{dockerExecutablePath, "rm", containerName}

	dockerCommand := &exec.Cmd{
		Path:   dockerExecutablePath,
		Args:   dockerRemoveContainerCommand,
		Stdout: loggerWriter,
		Stderr: loggerWriter,
	}

	dockerCommand.Run()

	logger.Println("Removing Image " + imageName)

	dockerRemoveImageCommand := []string{dockerExecutablePath, "rmi", imageName}

	dockerCommand = &exec.Cmd{
		Path:   dockerExecutablePath,
		Args:   dockerRemoveImageCommand,
		Stdout: loggerWriter,
		Stderr: loggerWriter,
	}

	dockerCommand.Run()

	logger.Println("Running image " + imageName + " with container name " + containerName)

	dockerCommand = &exec.Cmd{
		Path:   dockerExecutablePath,
		Args:   []string{dockerExecutablePath, "run", "-d", "-p", portMapping, "--name", containerName, imageName},
		Stdout: loggerWriter,
		Stderr: loggerWriter,
	}

	err := dockerCommand.Run()

	if err != nil {
		logger.Println(err.Error())
	}

	return err
}
