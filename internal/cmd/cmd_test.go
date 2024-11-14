package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetKeyName(t *testing.T) {
	key, err := GetKeyName(nil)
	assert.Error(t, err)
	assert.Empty(t, key)

	key, err = GetKeyName([]string{"test"})
	assert.NoError(t, err)
	assert.Equal(t, "test", key)
}
