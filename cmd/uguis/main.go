package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/yosssi/uguis"
)

const serviceNameMain = "main"

func main() {
	// Parse the command-line flags.
	showVersion := flag.Bool("v", false, "show version")
	showHelp := flag.Bool("h", false, "show help")
	twitterAPIKey := flag.String("twitter-api-key", "", "Twitter API key")
	twitterAPISecret := flag.String("twitter-api-secret", "", "Twitter API secret")
	twitterAccessToken := flag.String("twitter-access-token", "", "Twitter access token")
	twitterAccessTokenSecret := flag.String("twitter-access-token-secret", "", "Twitter access token secret")

	flag.Parse()

	if *showHelp {
		flag.PrintDefaults()
		return
	}

	if *showVersion {
		fmt.Println(uguis.Version)
		return
	}

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

	// Create a twitter client.
	twitterClient := uguis.NewSimpleTwitterClient(
		*twitterAPIKey,
		*twitterAPISecret,
		*twitterAccessToken,
		*twitterAccessTokenSecret,
		app,
		lgr,
		nil,
	)

	fmt.Println(twitterClient)
}
