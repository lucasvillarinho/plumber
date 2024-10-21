package cmd

import (
	"fmt"

	configs "github.com/lucasvillarinho/plumber/config"
	inj "github.com/lucasvillarinho/plumber/internal/injector"
)

type Cmd struct {
	plumberConfig *configs.PlumberConfig
}

func NewCmd(injector *inj.Injector) (*Cmd, error) {
	plumberConfig, err := inj.Get[*configs.PlumberConfig](injector)
	if err != nil || plumberConfig == nil {
		return nil, fmt.Errorf("error to get plumber config: %w", err)
	}

	return &Cmd{
		plumberConfig: *plumberConfig,
	}, nil
}

func (cmd *Cmd) Run() {

}
