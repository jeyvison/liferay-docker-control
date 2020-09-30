package main

import (
	"errors"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"github.com/jeyvison/liferay-docker-control/ldcDocker"
)

var dockerControl = ldcDocker.DockerControl{}

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

		var err error = nil

		switch liferayVersionRadio.Selected {
		case "Liferay CE":
			dockerControl.StopContainer("liferay-dxp-master")
			err = dockerControl.RunDocker("liferay-master", "jeyvison/liferay-master:latest", "8080:8080")
		case "Liferay DXP":
			dockerControl.StopContainer("liferay-master")
			err = dockerControl.RunDocker("liferay-dxp-master", "192.168.109.41:5000/jeyvison/liferay-dxp-master:latest", "8081:7300")
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
