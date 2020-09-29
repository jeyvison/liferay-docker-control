package main

import (
	"errors"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"github.com/mitchellh/go-homedir"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func stopContainer(containerName string, logger *log.Logger, stdWriter io.Writer) {
	dockerExecutablePath := getDockerExecutablePath()

	dockerStopContainerCommand := []string{dockerExecutablePath, "stop", containerName}

	logger.Println("Stopping container " + containerName)

	dockerCommand := &exec.Cmd{
		Path:   dockerExecutablePath,
		Args:   dockerStopContainerCommand,
		Stdout: stdWriter,
		Stderr: stdWriter,
	}

	dockerCommand.Run()
}

func runDocker(containerName string, imageName string, logger *log.Logger, portMapping string, stdWriter io.Writer) error {

	dockerExecutablePath := getDockerExecutablePath()

	logger.Println("Stopping container " + containerName)

	stopContainer(containerName, logger, stdWriter)

	logger.Println("Removing container " + containerName)

	dockerRemoveContainerCommand := []string{dockerExecutablePath, "rm", containerName}

	dockerCommand := &exec.Cmd{
		Path:   dockerExecutablePath,
		Args:   dockerRemoveContainerCommand,
		Stdout: stdWriter,
		Stderr: stdWriter,
	}

	dockerCommand.Run()

	logger.Println("Removing Image " + imageName)

	dockerRemoveImageCommand := []string{dockerExecutablePath, "rmi", imageName}

	dockerCommand = &exec.Cmd{
		Path:   dockerExecutablePath,
		Args:   dockerRemoveImageCommand,
		Stdout: stdWriter,
		Stderr: stdWriter,
	}

	dockerCommand.Run()

	logger.Println("Running image " + imageName + " with container name " + containerName)

	dockerCommand = &exec.Cmd{
		Path:   dockerExecutablePath,
		Args:   []string{dockerExecutablePath, "run", "-d", "-p", portMapping, "--name", containerName, imageName},
		Stdout: stdWriter,
		Stderr: stdWriter,
	}

	err := dockerCommand.Run()

	if err != nil {
		logger.Println(err.Error())
	}

	return err
}

func getDockerExecutablePath() string {
	return "/usr/local/bin/docker"
}

func imageVersions() (*widget.Radio, *widget.Box) {

	radio := widget.NewRadio([]string{"Liferay CE", "Liferay DXP"}, nil)

	radio.Horizontal = true

	hbox := widget.NewHBox(radio)

	return radio, hbox
}

func main() {
	a := app.New()
	w := a.NewWindow("Liferay Docker Control")

	w.Resize(fyne.NewSize(300, 300))

	vbox := widget.NewVBox()

	liferayVersionRadio, liferayVersionRadioBox := imageVersions()

	liferayVersionRadio.Required = true

	vbox.Append(liferayVersionRadioBox)

	button := widget.NewButton("Create/Update Liferay", nil)

	vbox.Append(button)

	progressBarInfinite := widget.NewProgressBarInfinite()

	progressBarInfinite.Hide()

	vbox.Append(progressBarInfinite)

	button.OnTapped = func() {
		progressBarInfinite.Show()
		button.Disable()

		logger, stdWriter := getLogger(vbox)

		var err error = nil

		switch liferayVersionRadio.Selected {
		case "Liferay CE":
			stopContainer("liferay-dxp-master", logger, stdWriter)
			err = runDocker("liferay-master", "jeyvison/liferay-master:latest", logger, "8080:8080", stdWriter)
		case "Liferay DXP":
			stopContainer("liferay-master", logger, stdWriter)
			err = runDocker("liferay-dxp-master", "192.168.109.41:5000/jeyvison/liferay-dxp-master:latest", logger, "8081:7300", stdWriter)
		default:
			err = errors.New("You must select one of of the versions")
		}

		if err != nil {
			vbox.Append(widget.NewLabel(err.Error()))
		}

		progressBarInfinite.Hide()
		button.Enable()
	}

	w.SetContent(vbox)

	w.ShowAndRun()
}

func getLogger(vbox *widget.Box) (*log.Logger, io.Writer) {
	var userDir, err = homedir.Dir()

	if err != nil {
		vbox.Append(widget.NewLabel(err.Error()))
	}

	logFilePath := filepath.FromSlash(userDir + "/liferay-docker-control.log")

	var LogFile, errFile = os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	if errFile != nil {
		vbox.Append(widget.NewLabel(errFile.Error()))
	}

	var logger = log.New(LogFile, "logging: ", log.Ldate)

	stdWriter := logger.Writer()
	return logger, stdWriter
}
