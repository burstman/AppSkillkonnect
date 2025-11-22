//go:generate fyne bundle -o bundle.go assets

package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	skilltheme "skillKonnectApp/pkg/theme"
	"skillKonnectApp/pkg/ui"
)

func main() {
	// Create the app
	a := app.New()

	// Use bundled theme icons from bundle.go
	lightIcon := resourceThemeLightlPng
	darkIcon := resourceDarckThemePng

	// Usage:
	a.Settings().SetTheme(skilltheme.NewSkillKonnectTheme(theme.VariantLight))

	// Main window
	w := a.NewWindow("SkillKonnect")
	w.SetMaster() // Uncomment this for full mobile app behavior (full screen, no window borders)

	// Simulate phone size on desktop (ignored on real phone)
	// Common Android phone sizes:
	// - iPhone SE: 375x667
	// - iPhone 12/13: 390x844
	// - Samsung Galaxy S21: 360x800
	// - Pixel 5: 393x851
	// - Average Android phone: ~360-410 width, ~800-900 height
	w.Resize(fyne.NewSize(390, 844)) // iPhone 12/13 size - good Android approximation

	// Theme toggle setup
	isDarkTheme := false // Start with LIGHT theme

	var ThemeBtn *widget.Button
	ThemeBtn = widget.NewButtonWithIcon("", darkIcon, func() {
		isDarkTheme = !isDarkTheme
		fmt.Println("Theme toggled. isDarkTheme:", isDarkTheme)

		if isDarkTheme {
			ThemeBtn.SetIcon(lightIcon)
			a.Settings().SetTheme(skilltheme.NewSkillKonnectTheme(theme.VariantDark))
			fmt.Println("Applied Dark theme")
		} else {
			ThemeBtn.SetIcon(darkIcon)
			a.Settings().SetTheme(skilltheme.NewSkillKonnectTheme(theme.VariantLight))
			fmt.Println("Applied Light theme")
		}
		w.Content().Refresh()
	})

	// Top bar with theme toggle button
	topBar := container.NewHBox(
		layout.NewSpacer(),
		ThemeBtn,
	)

	// Welcome screen with "Get Started" button that opens main screen
	welcomeScreen := ui.CreateWelcomeScreen(func() {
		// When "Get Started" is clicked, show the main screen
		mainScreen := ui.CreateMainScreen()
		mainLayout := container.NewBorder(
			topBar,     // Top (keep theme button)
			nil,        // Bottom
			nil,        // Left
			nil,        // Right
			mainScreen, // Center (main screen content)
		)
		w.SetContent(mainLayout)
	})

	// Initial layout with welcome screen
	initialLayout := container.NewBorder(
		topBar,        // Top
		nil,           // Bottom
		nil,           // Left
		nil,           // Right
		welcomeScreen, // Center (welcome screen)
	)

	w.SetContent(initialLayout)

	// Make sure window is visible
	w.Show()

	// Center the window on screen
	w.CenterOnScreen()

	// Show and run
	w.ShowAndRun()
}
