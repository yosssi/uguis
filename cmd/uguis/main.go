package main

import "github.com/yosssi/uguis"

const serviceNameMain = "main"

func main() {
	// Create a logger.
	lgr := uguis.NewSimpleLogger(nil)

	// Defer the call of the close method.
	defer func() {
		if err := lgr.Close(); err != nil {
			panic(err)
		}
	}()

	lgr.Print(uguis.NewLog(
		uguis.LogLevelINFO,
		"",
		serviceNameMain,
		"test",
	))
}
