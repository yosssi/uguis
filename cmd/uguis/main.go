package main

import (
	"fmt"
	"runtime"

	"github.com/yosssi/uguis"
)

const serviceNameMain = "main"

func main() {
	// Create an application.
	// This does nothing but holds an application configuration.
	app := uguis.NewApplication(nil)

	// Create a logger.
	lgr := uguis.NewSimpleLogger(nil)

	// Defer the call of the logger's close method.
	defer func() {
		if err := lgr.Close(); err != nil {
			panic(err)
		}
	}()

	// Set the maximum number of CPUs.
	runtime.GOMAXPROCS(app.CPUs)

	lgr.Print(uguis.NewLog(
		uguis.LogLevelINFO,
		app.Hostname,
		serviceNameMain,
		fmt.Sprintf("The maximum number of CPUs was set to %d.", runtime.GOMAXPROCS(0)),
	))
}
