package test

import (
	"golangplayground/testing/module"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestAdd(t *testing.T) {
	t.Run("must success ", func(t *testing.T) {
		assert.Equal(t, module.Add(1, 2), 3)
	})
}
