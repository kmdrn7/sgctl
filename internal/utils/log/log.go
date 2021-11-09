package log

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func Setup() {

	// Set generated log format
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "Mon, 02 Jan 2006 15:04:05",
	})

	// Output log to stdout instead of file
	log.SetOutput(os.Stdout)
}
