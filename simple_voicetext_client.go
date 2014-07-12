package uguis

import "github.com/yosssi/go-voicetext"

const serviceNameSimpleVoicetextClient = "simpleVoicetextClient"

// simpleVoicetextClient represents a simple voicetext client.
type simpleVoicetextClient struct {
	voicetext.Client
	reqC       chan VoicetextTTSRequest
	resC       chan *voicetext.Result
	closedReqC chan struct{}
	app        *Application
	lgr        Logger
}

// TTS calls the Voicetext TTS API.
func (c *simpleVoicetextClient) TTS(req VoicetextTTSRequest) {
	// Send a request to the tts goroutine.
	c.reqC <- req
}

// Close closes the voicetext client.
func (c *simpleVoicetextClient) Close() error {
	// Close the request channle.
	close(c.reqC)

	// Wait until the call goroutine is closed.
	<-c.closedReqC

	return nil
}

// tts calls a voicetext API.
func (c *simpleVoicetextClient) tts() {
	for req := range c.reqC {
		result, err := c.Client.TTS(req.text, req.opts)

		if err != nil {
			c.lgr.Print(NewLog(
				LogLevelERROR,
				c.app.Hostname,
				serviceNameSimpleVoicetextClient,
				err.Error(),
			))

			continue
		}

		c.resC <- result
	}

	// Send a closed signal.
	c.closedReqC <- struct{}{}
}

// NewSimpleVoicetextClient creates and returns a simple voicetext client.
func NewSimpleVoicetextClient(
	apiKey string,
	app *Application,
	lgr Logger,
	opts *SimpleVoicetextClientOptions,
) VoicetextClient {
	// Initialize options.
	if opts == nil {
		opts = &SimpleVoicetextClientOptions{}
	}
	opts.setDefaults()

	c := &simpleVoicetextClient{
		Client:     voicetext.NewClient(apiKey, nil),
		reqC:       make(chan VoicetextTTSRequest, opts.ReqCBfSize),
		resC:       make(chan *voicetext.Result, opts.ResCBfSize),
		closedReqC: make(chan struct{}),
		app:        app,
		lgr:        lgr,
	}

	// Launch a call goroutine.
	go c.tts()

	return c
}
