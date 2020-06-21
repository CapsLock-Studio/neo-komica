package config

import (
	"github.com/CapsLock-Studio/neo-komica/util"
	"github.com/golang/glog"
)

// Sentry - The config of tappay app
type Sentry struct {
	DSN         string
	Environment string
}

// NewSentry - get default sentry config
func NewSentry() *Sentry {
	sentry := &Sentry{}
	sentry.DSN = util.Getenv("SENTRY_DSN", "")
	sentry.Environment = util.Getenv("SENTRY_ENVIRONMENT", "development")

	glog.Infof("Sentry DSN: %v", sentry.DSN)

	return sentry
}
