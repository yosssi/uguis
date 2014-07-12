package uguis

import (
	"fmt"
	"os/exec"
)

const serviceNameSimplePlayer = "simplePlayer"

// simplePlayer represents a simple player.
type simplePlayer struct {
	command    string
	reqC       chan PlayerRequest
	resC       chan PlayerResponse
	closedReqC chan struct{}
	app        *Application
	lgr        Logger
}

// Play plays a sound file.
func (p *simplePlayer) Play(req PlayerRequest) {
	// Send a request to the play goroutine.
	p.reqC <- req
}

// Close closes the player.
func (p *simplePlayer) Close() error {
	// Close the request channel.
	close(p.reqC)

	// Wait until the play goroutine is closed.
	<-p.closedReqC

	return nil
}

// ResC returns a response channel.
func (p *simplePlayer) ResC() <-chan PlayerResponse {
	return p.resC
}

// play plays a sound file.
func (p *simplePlayer) play() {
	for req := range p.reqC {
		p.lgr.Print(NewLog(
			LogLevelINFO,
			p.app.Hostname,
			serviceNameSimplePlayer,
			fmt.Sprintf("%s by %s(@%s)", req.tweet.Text, req.tweet.User.Name, req.tweet.User.ScreenName),
		))

		path := req.path

		if err := exec.Command(p.command, path).Run(); err != nil {
			p.logError(err)
			continue
		}
		p.resC <- NewPlayerResponse(req.tweet, path)
	}

	// Send a closed signal.
	p.closedReqC <- struct{}{}
}

func (p *simplePlayer) logError(err error) {
	p.lgr.Print(NewLog(
		LogLevelERROR,
		p.app.Hostname,
		serviceNameSimplePlayer,
		err.Error(),
	))
}

// NewSimplePlayer creates and returns a simple player.
func NewSimplePlayer(
	command string,
	app *Application,
	lgr Logger,
	opts *SimplePlayerOptions,
) Player {
	// Initialize options.
	if opts == nil {
		opts = &SimplePlayerOptions{}
	}
	opts.setDefaults()

	p := &simplePlayer{
		command:    command,
		reqC:       make(chan PlayerRequest, opts.ReqCBfSize),
		resC:       make(chan PlayerResponse, opts.ResCBfSize),
		closedReqC: make(chan struct{}),
		app:        app,
		lgr:        lgr,
	}

	// Launch a play goroutine.
	go p.play()

	return p
}
