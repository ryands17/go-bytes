package structures_test

import (
	"testing"

	"github.com/peterldowns/testy/assert"
	"github.com/ryands17/go-bytes/cmd/structures"
)

func TestSetOperations(t *testing.T) {
	s := structures.NewSet[int]()
	t.Run("Set can add elements successfully", func(t *testing.T) {
		s.Add(1)
		s.Add(2)
		s.Add(3)
		s.Add(3)
		s.Add(3)

		assert.Equal(t, 3, s.Size())
		assert.True(t, s.Contains(1))
		assert.True(t, s.Contains(2))
		assert.True(t, s.Contains(3))
	})

	t.Run("Set can remove elements successfully", func(t *testing.T) {
		s.Remove(2)
		s.Remove(2)

		assert.Equal(t, 2, s.Size())
		assert.False(t, s.Contains(2))
	})
}
