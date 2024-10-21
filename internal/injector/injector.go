package pkg

import (
	"fmt"

	"github.com/samber/do"
)

type Injector struct {
	container *do.Injector
}

func NewInjector() (*Injector, error) {
	return &Injector{
		container: do.New(),
	}, nil
}

func Register[T any](injector *Injector, constructor func(*Injector) (T, error)) error {
	wrappedConstructor := func(di *do.Injector) (T, error) {
		return constructor(injector)
	}

	do.Provide(injector.container, wrappedConstructor)
	return nil
}

func Get[T any](injector *Injector) (*T, error) {
	dep, err := do.Invoke[T](injector.container)
	if err != nil {
		return nil, fmt.Errorf("error to get dependency: %w", err)
	}
	return &dep, nil
}
