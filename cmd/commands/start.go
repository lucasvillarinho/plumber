package commands

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/lucasvillarinho/plumber/internal/injector"
)

func Start(i *injector.Injector) *cobra.Command {

	return &cobra.Command{
		Use:   "start",
		Short: "Runs all processes",
		Long:  "The start command runs all processes defined in the Procfile concurrently",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Olá, seja bem-vindo!")
		},
	}

}
