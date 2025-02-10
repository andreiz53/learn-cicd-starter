package auth

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetAPIKey(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		header := http.Header{
			"Authorization": []string{"ApiKey asdasdasdasd"},
		}
		key, err := GetAPIKey(header)
		require.NoError(t, err)
		require.NotEmpty(t, key)
	})
	t.Run("Empty", func(t *testing.T) {
		header := http.Header{
			"Authorization": []string{""},
		}
		key, err := GetAPIKey(header)
		require.Error(t, err)
		require.Empty(t, key)
		require.Equal(t, err, ErrNoAuthHeaderIncluded)
	})
	t.Run("Invalid", func(t *testing.T) {
		header := http.Header{
			"Authorization": []string{"ApiKey"},
		}
		key, err := GetAPIKey(header)
		require.Error(t, err)
		require.Empty(t, key)
	})
}
