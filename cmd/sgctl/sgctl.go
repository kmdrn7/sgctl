package sgctl

import (
	"sgctl/internal/pkg/sgctl"
	"sgctl/internal/utils/config"
)

func Start() {
	// Check is running in dryrun mode or not
	if config.Config.Dryrun {
		// Run dryrun command
		sgctl.Dryrun()
	} else {
		// Run command
		sgctl.Run()
	}
}
