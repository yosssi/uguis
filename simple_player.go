package uguis

import "fmt"

const serviceNameSimplePlayer = "simplePlayer"

// simplePlayer represents a simple player.
type simplePlayer struct {
	reqC       chan string
	resC       chan string
	closedReqC chan struct{}
	app        *Application
	lgr        Logger
}

// Play plays a sound file.
func (p *simplePlayer) Play(path string) {
	// Send a request to the play goroutine.
	p.reqC <- path
}

// Close closes the player.
func (p *simplePlayer) Close() error {
	// Close the request channel.
	close(p.reqC)

	// Wait until the play goroutine is closed.
	<-p.closedReqC

	return nil
}

// play plays a sound file.
func (p *simplePlayer) play() {
	for path := range p.reqC {
		//TODO
		fmt.Println(path)
	}

	// Send a closed signal.
	p.closedReqC <- struct{}{}
}

// NewSimplePlayer creates and returns a simple player.
func NewSimplePlayer(
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
		reqC:       make(chan string, opts.ReqCBfSize),
		resC:       make(chan string, opts.ResCBfSize),
		closedReqC: make(chan struct{}),
		app:        app,
		lgr:        lgr,
	}

	// Launch a play goroutine.
	go p.play()

	return p
}
