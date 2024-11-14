package otp

import (
	"testing"

	"github.com/mojixcoder/otpgen/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestGenerateOTP_Error(t *testing.T) {
	code, err := GenerateOTP(config.TOTPConfig{})

	assert.Error(t, err)
	assert.Empty(t, code)
}

func TestGenerateOTP(t *testing.T) {
	code, err := GenerateOTP(config.TOTPConfig{Secret: "6JVT452ZKRWRY65LQTZLCPCNJUZDLH2G"})

	assert.NoError(t, err)
	assert.NotEmpty(t, code)
}
