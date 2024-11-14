package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mojixcoder/otpgen/internal/cmd"
	"github.com/mojixcoder/otpgen/internal/config"
	"github.com/mojixcoder/otpgen/internal/otp"
)

func checkErr(err error) {
	if err == nil {
		return
	}

	fmt.Println(err.Error() + ".")
	os.Exit(1)
}

func main() {
	flag.Parse()

	key, err := cmd.GetKeyName(flag.Args())
	checkErr(err)

	configs, err := config.ReadConfigs(*cmd.ConfigFlag)
	checkErr(err)

	totpCfg, ok := configs.Keys[key]
	if !ok {
		fmt.Printf("key(%s) was not found.\n", key)
		os.Exit(1)
	}

	code, err := otp.GenerateOTP(totpCfg)
	checkErr(err)

	fmt.Println(code)
	os.Exit(0)
}
