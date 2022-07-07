package test

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestAdd(t *testing.T) {
	t.Run("must success ", func(t *testing.T) {
		assert.Equal(t, module.add(1, 2), 3)

	})
}
