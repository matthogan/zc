package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/go-logr/logr"
	"github.com/matthogan/zc/cmd/cn/cli"
	logging "github.com/matthogan/zc/pkg/logging"
)

const (
	APP = "cn"
)

var (
	root logr.Logger
)

func init() {
	defer finally()
	root = logging.Init()
}

func main() {
	defer finally()
	root.Info(APP)
	for i, arg := range os.Args {
		if (strings.HasPrefix(arg, "-") && len(arg) == 2) || (strings.HasPrefix(arg, "--") && len(arg) >= 4) {
			continue
		}
		if strings.HasPrefix(arg, "--") && len(arg) == 3 {
			newArg := fmt.Sprintf("-%c", arg[2])
			fmt.Fprintf(os.Stderr, "WARNING: the flag %s is deprecated and will be removed in a future release. Please use the flag %s.\n", arg, newArg)
			os.Args[i] = newArg
		} else if strings.HasPrefix(arg, "-") {
			newArg := fmt.Sprintf("-%s", arg)
			fmt.Fprintf(os.Stderr, "WARNING: the flag %s is deprecated and will be removed in a future release. Please use the flag %s.\n", arg, newArg)
			os.Args[i] = newArg
		}
	}
	if err := cli.New().Execute(); err != nil {
		panic(err)
	}
}

// catch panics so stack traces are not output
func finally() {
	if err := recover(); err != nil {
		fmt.Fprintf(os.Stdout, "Error: %s\n", err)
	}
}
