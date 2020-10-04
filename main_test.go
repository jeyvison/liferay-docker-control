package main

import (
	"fyne.io/fyne/test"
	"fyne.io/fyne/widget"
	"github.com/akyoto/assert"
	"testing"
)

func TestDownloadWithoutSelectImageVersion(t *testing.T) {
	mainControl := newMainControl()
	mainControl.loadUI(test.NewApp())

	test.Tap(mainControl.buttons["Create/Update Liferay"])

	vbox := mainControl.vbox

	vboxChildrens := vbox.Children

	child := vboxChildrens[len(vboxChildrens)-1]

	ct, ok := child.(*widget.Label)

	assert.True(t, ok)
	assert.Equal(t, ct.Text, "You must select one of of the versions")
}
