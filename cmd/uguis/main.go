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
	voicetextAPIKey := flag.String("voicetext-api-key", "", "Voicetext Web API key")

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

	// Defer the call of the twitter client's close method.
	defer func() {
		if err := twitterClient.Close(); err != nil {
			panic(err)
		}
	}()

	// Create a voicetext client.
	voicetextClient := uguis.NewSimpleVoicetextClient(*voicetextAPIKey, app, lgr, nil)

	// Defer the call of the voicetext client's close method.
	defer func() {
		if err := voicetextClient.Close(); err != nil {
			panic(err)
		}
	}()

	// Create a file writer.
	fileWriter := uguis.NewSimpleFileWriter(app, lgr, nil)

	// Defer the call of the file writer's close method.
	defer func() {
		if err := fileWriter.Close(); err != nil {
			panic(err)
		}
	}()

	// Create a player.
	player := uguis.NewSimplePlayer(app, lgr, nil)

	// Defer the call of the player's close method.
	defer func() {
		if err := player.Close(); err != nil {
			panic(err)
		}
	}()

	// Create a controller.
	ctrl := uguis.NewController(app, lgr, twitterClient, voicetextClient, fileWriter, player)

	// Executes the controller's main process.
	ctrl.Exec()
}
