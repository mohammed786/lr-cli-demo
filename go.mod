module lr-cli

go 1.14

require (
	github.com/caarlos0/env v3.5.0+incompatible
	github.com/loginradius/lr-cli v0.0.0-00010101000000-000000000000
	github.com/mitchellh/go-homedir v1.1.0
	github.com/spf13/cobra v1.1.3
	github.com/spf13/viper v1.7.1
)

replace github.com/loginradius/lr-cli => ../cli
