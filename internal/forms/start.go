package forms

import (
	"fmt"

	"github.com/charmbracelet/huh"

	"github.com/lucasvillarinho/plumber/internal/fields"
	"github.com/lucasvillarinho/plumber/internal/injector"
)

func Start(i *injector.Injector) error {
	fmt.Println()

	if err := welcomeForm(); err != nil {
		return err
	}

	return nil
}

func welcomeForm() error {
	return huh.NewForm(huh.NewGroup(fields.WelcomeNote())).Run()
}
