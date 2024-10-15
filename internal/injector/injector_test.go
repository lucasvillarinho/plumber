package injector

import (
	"errors"
	"testing"

	"github.com/samber/do"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRegister(t *testing.T) {
	t.Run("Should return nil when constructor is valid ", func(t *testing.T) {
		constructor := func(inj *do.Injector) (string, error) {
			return "test-value", nil
		}

		injector := NewInjector()

		err := Register[string](injector, constructor)

		assert.NoError(t, err, "expected no error when constructor is valid")
	})

	t.Run("Should return error when constructor is invalid", func(t *testing.T) {
		invalidConstructor := func() {}

		injector := NewInjector()

		err := Register[string](injector, invalidConstructor)

		assert.Error(t, err, "Expected an error when constructor is invalid")
		assert.Equal(t, errors.New("constructor must be a function that takes a *do.Injector and returns a value and an error"), err, "Expected specific error message when constructor is invalid")
	})
}

func TestGet(t *testing.T) {
	t.Run("Should retrieve the dependency successfully", func(t *testing.T) {
		expected := "test-value"

		injector := NewInjector()

		do.Provide(injector.container, func(inj *do.Injector) (string, error) {
			return expected, nil
		})

		result, err := Get[string](injector)

		require.NoError(t, err, "Expected no error when dependency is resolved successfully")
		assert.Equal(t, &expected, result, "Expected result to be the expected value when dependency is resolved")
	})

	t.Run("Should fail to retrieve the dependency when it cannot be resolved", func(t *testing.T) {
		injector := NewInjector()

		do.Provide(injector.container, func(inj *do.Injector) (string, error) {
			return "", errors.New("dependency not found")
		})

		result, err := Get[string](injector)

		require.Error(t, err, "Expected an error when dependency cannot be resolved")
		assert.Nil(t, result, "Expected result to be nil when dependency resolution fails")

		assert.Contains(t, err.Error(), "failed to get dependency", "Expected error message to contain 'failed to get dependency'")
	})
}
