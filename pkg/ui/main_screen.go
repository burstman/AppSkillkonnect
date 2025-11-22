package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// CreateMainScreen creates the main application screen (shown after welcome)
func CreateMainScreen() fyne.CanvasObject {
	// Empty screen for now - you can add your main UI here later
	title := widget.NewLabel("Main Screen")
	title.Alignment = fyne.TextAlignCenter

	content := container.NewCenter(title)

	return content
}
