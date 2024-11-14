package config

import (
	"os"
	"testing"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"github.com/stretchr/testify/assert"
)

func TestReadConfigs_FileNotFound(t *testing.T) {
	cfg, err := ReadConfigs("invalid_file.yaml")

	assert.ErrorIs(t, err, os.ErrNotExist)
	assert.Empty(t, cfg)
}

func TestReadConfigs_InvalidYamlFile(t *testing.T) {
	file, err := os.CreateTemp(".", "config_*.yaml")
	if !assert.NoError(t, err) {
		return
	}
	defer file.Close()

	_, err = file.WriteString(`keys:
			- 2
	- 1
	`)
	if !assert.NoError(t, err) {
		return
	}

	cfg, err := ReadConfigs(file.Name())
	assert.Error(t, err)
	assert.Empty(t, cfg)

	assert.NoError(t, os.Remove(file.Name()))
}

func TestReadConfigs(t *testing.T) {
	file, err := os.CreateTemp(".", "config_*.yaml")
	if !assert.NoError(t, err) {
		return
	}
	defer file.Close()

	_, err = file.WriteString(`keys:
  test:
    secret: "1234"
`)
	if !assert.NoError(t, err) {
		return
	}

	cfg, err := ReadConfigs(file.Name())
	assert.NoError(t, err)
	assert.Equal(
		t,
		Config{
			Keys: map[string]TOTPConfig{"test": {
				Secret: "1234",
			}},
		},
		cfg,
	)

	assert.NoError(t, os.Remove(file.Name()))
}

func TestTOTPConfig_ToValidateOpts(t *testing.T) {
	cfg := TOTPConfig{
		Skew:      new(uint),
		Period:    30,
		Algorithm: otp.AlgorithmSHA1,
		Digits:    otp.DigitsSix,
	}

	assert.Equal(
		t,
		totp.ValidateOpts{
			Skew:      0,
			Period:    30,
			Algorithm: otp.AlgorithmSHA1,
			Digits:    otp.DigitsSix,
		},
		cfg.ToValidateOpts(),
	)
}

func TestIsValid_Error(t *testing.T) {
	cfg := TOTPConfig{}
	assert.Error(t, cfg.IsValid())
}

func TestIsValid(t *testing.T) {
	cfg := TOTPConfig{
		Secret: " abc ",
	}

	skew := uint(1)
	assert.NoError(t, cfg.IsValid())
	assert.Equal(
		t,
		TOTPConfig{
			Secret:    "abc",
			Digits:    otp.DigitsSix,
			Period:    30,
			Algorithm: otp.AlgorithmSHA1,
			Skew:      &skew,
		},
		cfg,
	)
}
