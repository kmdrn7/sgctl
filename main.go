package main

import (
	"sgctl/cmd/sgctl"
	"sgctl/internal/utils/config"
	"sgctl/internal/utils/log"
)

func main() {
	// Setup logging
	log.Setup()

	// Setup config
	config.Setup()

	// Start main program
	sgctl.Start()
}
