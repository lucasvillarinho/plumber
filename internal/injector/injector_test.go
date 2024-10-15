package injector

import (
	"errors"
	"testing"

	"github.com/samber/do"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {

	t.Run("Should return nil when constructor is invalid ", func(t *testing.T) {
		constructor := func(inj *do.Injector) (string, error) {
			return "test-value", nil
		}

		injector := NewInjector()

		err := Register[string](injector, constructor)

		assert.NoError(t, err)
		assert.Nil(t, err)
	})

	t.Run("Should return erro when constructor is invalid", func(t *testing.T) {
		invalidConstructor := func() {}

		injector := NewInjector()

		err := Register[string](injector, invalidConstructor)

		assert.Error(t, err)
		assert.Equal(t, errors.New("constructor must be a function that takes a *do.Injector and returns a value and an error"), err)
	})
}
