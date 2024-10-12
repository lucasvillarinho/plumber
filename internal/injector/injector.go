package injector

import (
	"errors"
	"fmt"

	"github.com/samber/do"
)

type Injector struct {
	container *do.Injector
}

func New() *Injector {
	return &Injector{
		container: do.New(),
	}
}

func Register[T any](injector *Injector, constructor any) error {
	if fn, ok := constructor.(func(*do.Injector) (T, error)); ok {
		do.Provide(injector.container, fn)
	}

	return errors.New("constructor must be a function that takes a *do.Injector and returns a value and an error")
}

func Get[T any](injector *Injector) (*T, error) {
	dep, err := do.Invoke[T](injector.container)
	if err != nil {
		return nil, fmt.Errorf("failed to get dependency: %w", err)
	}

	return &dep, nil
}
