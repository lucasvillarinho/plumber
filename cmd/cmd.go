package cmd

import (
	"github.com/spf13/cobra"

	"github.com/lucasvillarinho/plumber/cmd/commands"
	"github.com/lucasvillarinho/plumber/internal/forms"
	"github.com/lucasvillarinho/plumber/internal/injector"
)

var rootCmd = &cobra.Command{
	Use: "app",
	RunE: func(cmd *cobra.Command, args []string) error {
		injectorInstance := injector.NewInjector()

		if err := run(injectorInstance); err != nil {
			return err
		}

		return nil
	},
}

func run(i *injector.Injector) error {
	if err := forms.Start(i); err != nil {
		return err
	}

	return nil
}

func Setup(i *injector.Injector) (*cobra.Command, error) {

	rootCmd.AddCommand(commands.Start(i))

	return rootCmd, nil
}
