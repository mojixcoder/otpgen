package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"gopkg.in/yaml.v3"
)

func ReadConfigs(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("error in reading config file, error: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return Config{}, fmt.Errorf("error in unmarshaling configs, error: %w", err)
	}

	return cfg, nil
}

type Config struct {
	Keys map[string]TOTPConfig `yaml:"keys"`
}

type TOTPConfig struct {
	Skew      *uint         `yaml:"skew"`
	Period    uint          `yaml:"period"`
	Digits    otp.Digits    `yaml:"digits"`
	Algorithm otp.Algorithm `yaml:"algorithm"`
	Secret    string        `yaml:"secret"`
}

func (cfg TOTPConfig) ToValidateOpts() totp.ValidateOpts {
	return totp.ValidateOpts{
		Period:    cfg.Period,
		Digits:    cfg.Digits,
		Algorithm: cfg.Algorithm,
		Skew:      *cfg.Skew,
	}
}

func (cfg *TOTPConfig) IsValid() error {
	if secret := strings.TrimSpace(cfg.Secret); secret == "" {
		return fmt.Errorf("secret is not specified")
	} else {
		cfg.Secret = secret
	}

	if cfg.Period == 0 {
		cfg.Period = 30
	}

	if cfg.Skew == nil {
		cfg.Skew = new(uint)
		*cfg.Skew = 1
	}

	if cfg.Digits == 0 {
		cfg.Digits = otp.DigitsSix
	}

	if cfg.Algorithm == 0 {
		cfg.Algorithm = otp.AlgorithmSHA1
	}

	return nil
}
