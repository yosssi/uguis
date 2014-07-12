package uguis

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const serviceNameController = "Controller"

var tempDir = os.TempDir()

// Controller represents a controller.
type Controller struct {
	app                 *Application
	lgr                 Logger
	twitterClient       TwitterClient
	voicetextClient     VoicetextClient
	fileWriter          FileWriter
	player              Player
	twitterSinceID      string
	twitterSinceIDMutex *sync.RWMutex
}

// Exec executes the controller's main process.
func (ctrl *Controller) Exec() {
	// Listen the twitter client.
	go ctrl.listenTwitterClient()

	// Listen the voicetext client.
	go ctrl.listenVoicetextClient()

	// Listen the file writer client.
	go ctrl.listenFileWriterClient()

	// Listen the player
	go ctrl.listenPlayer()

	// Call the Twitter API.
	params := map[string]string{
		"count": "200",
	}
	for {
		if sinceID := ctrl.getTwitterSinceID(); sinceID != "" {
			params["since_id"] = sinceID
		}

		ctrl.twitterClient.Call(newTwitterRequest(
			httpGET,
			"/statuses/home_timeline.json",
			params,
		))

		time.Sleep(time.Minute)
	}
}

func (ctrl *Controller) listenTwitterClient() {
	for tweet := range ctrl.twitterClient.ResC() {
		ctrl.setTwitterSinceID(tweet.IDStr)

		// Call the Voicetext API.
		ctrl.voicetextClient.TTS(NewVoicetextTTSRequest(tweet, nil))
	}
}

func (ctrl *Controller) listenVoicetextClient() {
	for res := range ctrl.voicetextClient.ResC() {
		if res.result.ErrMsg != nil {
			ctrl.lgr.Print(NewLog(
				LogLevelERROR,
				ctrl.app.Hostname,
				serviceNameController,
				res.result.ErrMsg.String(),
			))
			continue
		}

		// Call the file writer client.
		ctrl.fileWriter.Write(NewFileWriterRequest(res.tweet, newFile(
			filepath.Join(tempDir, fmt.Sprintf(filenameFormat, time.Now().UnixNano())),
			res.result.Sound,
			fileChangeTypeCreate,
		)))
	}
}

func (ctrl *Controller) listenFileWriterClient() {
	for res := range ctrl.fileWriter.ResC() {
		// Call the player.
		ctrl.player.Play(NewPlayerRequest(res.tweet, res.path))
	}
}

func (ctrl *Controller) listenPlayer() {
	for res := range ctrl.player.ResC() {
		// Call the file writer client to delete the temporary file.
		ctrl.fileWriter.Write(NewFileWriterRequest(res.tweet, newFile(
			res.path,
			nil,
			fileChangeTypeDelete,
		)))
	}
}

func (ctrl *Controller) getTwitterSinceID() string {
	ctrl.twitterSinceIDMutex.RLock()
	id := ctrl.twitterSinceID
	ctrl.twitterSinceIDMutex.RUnlock()
	return id
}

func (ctrl *Controller) setTwitterSinceID(id string) {
	ctrl.twitterSinceIDMutex.Lock()
	ctrl.twitterSinceID = id
	ctrl.twitterSinceIDMutex.Unlock()
}

// NewController creates and returns a controller.
func NewController(
	app *Application,
	lgr Logger,
	twitterClient TwitterClient,
	voicetextClient VoicetextClient,
	fileWriter FileWriter,
	player Player,
) *Controller {
	return &Controller{
		app:                 app,
		lgr:                 lgr,
		twitterClient:       twitterClient,
		voicetextClient:     voicetextClient,
		fileWriter:          fileWriter,
		player:              player,
		twitterSinceIDMutex: new(sync.RWMutex),
	}
}
