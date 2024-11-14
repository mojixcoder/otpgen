package otp

import (
	"time"

	"github.com/mojixcoder/otpgen/internal/config"
	"github.com/pquerna/otp/totp"
)

func GenerateOTP(cfg config.TOTPConfig) (string, error) {
	if err := cfg.IsValid(); err != nil {
		return "", err
	}

	return totp.GenerateCodeCustom(cfg.Secret, time.Now().UTC(), cfg.ToValidateOpts())
}
