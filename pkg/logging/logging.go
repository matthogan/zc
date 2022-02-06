package logging

import (
	"flag"

	"github.com/go-logr/logr"
	"go.uber.org/zap/zapcore"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

func Init() logr.Logger {
	opts := zap.Options{
		Development: true,
		TimeEncoder: zapcore.ISO8601TimeEncoder,
	}
	opts.BindFlags(flag.CommandLine)
	ctrl.SetLogger(zap.New(zap.UseFlagOptions(&opts)))
	return ctrl.Log
}

// Returns a k8s logr if running in k8s, otherwise returns a simple stdout logger
func Logger(name string) logr.Logger {
	return ctrl.Log.WithName(name)
}
