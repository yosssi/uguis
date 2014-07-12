package uguis

const serviceNameController = "Controller"

// Controller represents a controller.
type Controller struct {
	app             *Application
	lgr             Logger
	twitterClient   TwitterClient
	voicetextClient VoicetextClient
	fileWriter      FileWriter
	player          Player
}

// Exec executes the controller's main process.
func (ctrl *Controller) Exec() {
	ctrl.lgr.Print(NewLog(
		LogLevelINFO,
		ctrl.app.Hostname,
		serviceNameController,
		"Exec!",
	))
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
		app:             app,
		lgr:             lgr,
		twitterClient:   twitterClient,
		voicetextClient: voicetextClient,
		fileWriter:      fileWriter,
		player:          player,
	}
}
