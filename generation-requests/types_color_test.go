package genreq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MatchColor(t *testing.T) {
	t.Run("basic tests", func(t *testing.T) {
		color, remainder, err := MatchColor("red")
		assert.NoError(t, err)
		assert.Equal(t, ColorRed, color)
		assert.Equal(t, "", remainder)

		color, remainder, err = MatchColor("red shoes")
		assert.NoError(t, err)
		assert.Equal(t, ColorRed, color)
		assert.Equal(t, "shoes", remainder)

		color, remainder, err = MatchColor("Green Muscadine grapes")
		assert.NoError(t, err)
		assert.Equal(t, ColorGreen, color)
		assert.Equal(t, "Muscadine grapes", remainder)

		color, remainder, err = MatchColor("red-orange")
		assert.NoError(t, err)
		assert.Equal(t, ColorRedOrange, color)
		assert.Equal(t, "", remainder)

		color, remainder, err = MatchColor("orange-red")
		assert.NoError(t, err)
		assert.Equal(t, ColorRedOrange, color)
		assert.Equal(t, "", remainder)

		color, remainder, err = MatchColor("orange")
		assert.NoError(t, err)
		assert.Equal(t, ColorOrange, color)
		assert.Equal(t, "", remainder)

		_, _, err = MatchColor("green-orange")
		assert.ErrorIs(t, err, ErrNoColor)
	})
	t.Run("all color constants match as valid colors", func(t *testing.T) {
		for _, color := range Colors {
			got, remainder, err := MatchColor(string(color))
			assert.NoError(t, err)
			assert.Equal(t, color, got)
			assert.Equal(t, "", remainder)
		}
	})
}

func Test_resolveLookupKey(t *testing.T) {
	assert.Equal(t, "red", resolveLookupKey("red", ""))
	assert.Equal(t, "red", resolveLookupKey("red", "red"))
	assert.Equal(t, "orange-red", resolveLookupKey("red", "orange"))
	assert.Equal(t, "orange-red", resolveLookupKey("orange", "red"))
	assert.Equal(t, "yellow", resolveLookupKey("Yellow", "yelLOW"))
}
