package main

import (
	"fmt"
	"os"
	"os/exec"

	"fyne.io/fyne"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func runDocker() {

	docker_executable := "/usr/local/bin/docker"

	docker_stop_container_commnad := []string{docker_executable, "stop", "liferay-master"}

	docker_command := &exec.Cmd{
		Path:   docker_executable,
		Args:   docker_stop_container_commnad,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	fmt.Println(docker_command.String())

	docker_command.Run()

	docker_remove_container_commnad := []string{docker_executable, "rm", "liferay-master"}

	docker_command = &exec.Cmd{
		Path:   docker_executable,
		Args:   docker_remove_container_commnad,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	fmt.Println(docker_command.String())

	docker_command.Run()

	docker_remove_image_commnad := []string{docker_executable, "rmi", "jeyvison/liferay-master:latest"}

	docker_command = &exec.Cmd{
		Path:   docker_executable,
		Args:   docker_remove_image_commnad,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	fmt.Println(docker_command.String())

	docker_command.Run()

	docker_command = &exec.Cmd{
		Path:   docker_executable,
		Args:   []string{docker_executable, "run", "-d", "-p", "8080:8080", "--name", "liferay-master", "jeyvison/liferay-master:latest"},
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	fmt.Println(docker_command.String())

	error := docker_command.Run()

	if error != nil {
		fmt.Println(error)
	}
}

func main() {
	a := app.New()
	w := a.NewWindow("Liferay Docker Control")

	w.Resize(fyne.NewSize(300, 300))

	button := widget.NewButton("Create/Update Liferay", nil)

	vbox := widget.NewVBox(button)

	ipb := widget.NewProgressBarInfinite()

	ipb.Hide()

	vbox.Append(ipb)

	button.OnTapped = func() {
		ipb.Show()
		button.Disable()
		runDocker()
		ipb.Hide()
		button.Enable()
	}

	w.SetContent(vbox)

	w.ShowAndRun()
}
