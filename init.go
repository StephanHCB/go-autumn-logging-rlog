package go_autumn_logging_rlog

import (
	aulogging "github.com/StephanHCB/go-autumn-logging"
	"github.com/StephanHCB/go-autumn-logging-rlog/implementation/logging"
)

func init() {
	aulogging.Logger = &logging.RlogLoggingImplementation{}
}
