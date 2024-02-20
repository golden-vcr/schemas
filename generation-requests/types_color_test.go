package genreq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MatchColor(t *testing.T) {
	t.Run("basic tests", func(t *testing.T) {
		color, err := MatchColor("red")
		assert.NoError(t, err)
		assert.Equal(t, ColorRed, color)

		color, err = MatchColor("red-orange")
		assert.NoError(t, err)
		assert.Equal(t, ColorRedOrange, color)

		color, err = MatchColor("orange-red")
		assert.NoError(t, err)
		assert.Equal(t, ColorRedOrange, color)

		color, err = MatchColor("orange")
		assert.NoError(t, err)
		assert.Equal(t, ColorOrange, color)

		_, err = MatchColor("green-orange")
		assert.ErrorIs(t, err, ErrNoColor)
	})
	t.Run("all color constants match as valid colors", func(t *testing.T) {
		for _, color := range Colors {
			got, err := MatchColor(string(color))
			assert.NoError(t, err)
			assert.Equal(t, color, got)
		}
	})
}

func Test_resolveLookupKey(t *testing.T) {
	assert.Equal(t, "red", resolveLookupKey("red", ""))
	assert.Equal(t, "red", resolveLookupKey("red", "red"))
	assert.Equal(t, "orange-red", resolveLookupKey("red", "orange"))
	assert.Equal(t, "orange-red", resolveLookupKey("orange", "red"))
}
