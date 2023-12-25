package app_test

import (
	"testing"

	"github.com/Anwert/fyne/v2/internal/app"
	"github.com/Anwert/fyne/v2/test"
)

func TestApplySettings_BeforeContentSet(t *testing.T) {
	a := test.NewApp()
	_ = a.NewWindow("NoContent")

	app.ApplySettings(a.Settings(), a)
}
