package config

import (
	flag "github.com/spf13/pflag"
)

type ConfigType struct {
	Dryrun bool
}

var Config ConfigType

func Setup() {
	// Setup flag args
	flag.BoolVar(&Config.Dryrun, "dryrun", false, "run heimdal-cli with dryrun mode")

	// Parse all flags
	flag.Parse()
}
