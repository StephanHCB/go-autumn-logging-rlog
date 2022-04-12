package leveledlogging

import (
	"context"
	"fmt"
	aulogging "github.com/StephanHCB/go-autumn-logging"
	auloggingapi "github.com/StephanHCB/go-autumn-logging/api"
	"github.com/go-logr/logr"
	_ "github.com/go-logr/logr"
)

type RlogLeveledLoggingImplementation struct {
	Ctx    context.Context
	Level  int
	Err    error
	Values map[string]string
}

func (lv *RlogLeveledLoggingImplementation) WithErr(err error) auloggingapi.LeveledLoggingImplementation {
	lv.Err = err
	return lv
}

func (lv *RlogLeveledLoggingImplementation) With(key string, value string) auloggingapi.LeveledLoggingImplementation {
	lv.Values[key] = value
	return lv
}

func (lv *RlogLeveledLoggingImplementation) Print(v ...interface{}) {
	rlogLogger := logr.FromContextOrDiscard(lv.Ctx).V(lv.Level)
	for k, v := range lv.Values {
		rlogLogger = rlogLogger.WithValues(k, v)
	}
	message := fmt.Sprint(v...)
	lv.callback(message)
	if lv.Err != nil {
		rlogLogger.Error(lv.Err, message)
	} else {
		rlogLogger.Info(message)
	}
}

func (lv *RlogLeveledLoggingImplementation) Printf(format string, v ...interface{}) {
	rlogLogger := logr.FromContextOrDiscard(lv.Ctx).V(lv.Level)
	for k, v := range lv.Values {
		rlogLogger = rlogLogger.WithValues(k, v)
	}
	message := fmt.Sprintf(format, v...)
	lv.callback(message)
	if lv.Err != nil {
		rlogLogger.Error(lv.Err, message)
	} else {
		rlogLogger.Info(message)
	}
}

var levelNames = []string{"FATAL", "ERROR", "WARN", "INFO", "DEBUG", "TRACE"}

func (lv *RlogLeveledLoggingImplementation) callback(message string) {
	if aulogging.LogEventCallback != nil {
		ctx := lv.Ctx
		if ctx == nil {
			ctx = context.Background()
		}
		aulogging.LogEventCallback(ctx, levelNames[lv.Level], message, lv.Err, lv.Values)
	}
}
