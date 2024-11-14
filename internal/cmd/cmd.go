package cmd

import (
	"flag"
	"fmt"
)

var ConfigFlag *string

func init() {
	ConfigFlag = flag.String("config", "/etc/otpgen/config.yaml", "path to yaml config file")
}

func GetKeyName(args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("unexpected number of arguments, expected 1 argument, got %d", len(args))
	}

	return args[0], nil
}
