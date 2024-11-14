
# otpgen

In the company where I work at, I need OTP code to login almost everywhere.  
I was tired of the Google Authenticator app, so I decided to write this small project to be able to generate OTPs on my laptop without having to check my phone.


## Installing

Clone the project

```bash
  git clone https://github.com/mojixcoder/otpgen
```

Go to the project directory

```bash
  cd otpgen
```

Build the project

```bash
  go build
```

Make it executable

```bash
  sudo mv ./otpgen /usr/local/bin    
  sudo chmod +x /usr/local/bin/otpgen
```

Place your configs at `/etc/otpgen/config.yaml`

```yaml
keys:
  gitlab:
    secret: "totp_secret_for_gitlab"
  vpn:
    secret: "totp_secret_for_vpn"
```

Generate an OTP for your VPN
```bash
$ otpgen vpn
123456
```

## Exporting secrets from Google Authenticator 

To export TOTP secrets from Google Authenticator, please follow the instructions in this [repo](https://github.com/krissrex/google-authenticator-exporter).
