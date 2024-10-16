package main

import (
	"fmt"
	"os"

	"golang.org/x/exp/slog"

	"github.com/lucasvillarinho/plumber/cmd"
	"github.com/lucasvillarinho/plumber/internal/injector"
)

func main() {
	injector := injector.NewInjector()

	rootCmd, err := cmd.Setup(injector)
	if err != nil {
		slog.Error("error setting up command: %w", err)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
