package module

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestAdd(t *testing.T) {
	t.Run("must success ", func(t *testing.T) {
		assert.NotEqual(t, Add(1, 2), 4)
	})
}
