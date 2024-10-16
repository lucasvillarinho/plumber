package fields

import "github.com/charmbracelet/huh"

func WelcomeNote() *huh.Note {
	return huh.
		NewNote().
		Title("🛁 Plumber").
		Description("Easily manage and run multiple processes in a Go development environment")
}
