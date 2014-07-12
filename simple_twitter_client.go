package uguis

import (
	"fmt"
	"net/http"

	"github.com/mrjones/oauth"
)

// Twitter API URLs
const (
	twitterRequestTokenURL   = "http://api.twitter.com/oauth/request_token"
	twitterAuthorizeTokenURL = "https://api.twitter.com/oauth/authorize"
	twitterAccessTokenURL    = "https://api.twitter.com/oauth/access_token"
)

const serviceNameSimpleTwitterClient = "simpleTwitterClient"

// simpleTwitterClient represents a simple Twitter client.
type simpleTwitterClient struct {
	consumer    *oauth.Consumer
	accessToken *oauth.AccessToken
	reqC        chan twitterRequest
	resC        chan TwitterResponse
	closedReqC  chan struct{}
	app         *Application
	lgr         Logger
}

// Call calls a Twitter API.
func (c *simpleTwitterClient) Call(req twitterRequest) {
	// Send a request to the call goroutine.
	c.reqC <- req
}

// Close closes the twitter client.
func (c *simpleTwitterClient) Close() error {
	// Close the request channel.
	close(c.reqC)

	// Wait until the call goroutine is closed.
	<-c.closedReqC

	return nil
}

// call calls a Twitter API.
func (c *simpleTwitterClient) call() {
	for req := range c.reqC {
		var httpRes *http.Response
		var err error

		switch req.method {
		case httpGET:
			httpRes, err = c.consumer.Get(req.url, req.params, c.accessToken)
		case httpPOST:
			httpRes, err = c.consumer.Post(req.url, req.params, c.accessToken)
		case httpDELETE:
			httpRes, err = c.consumer.Delete(req.url, req.params, c.accessToken)
		default:
			err = fmt.Errorf("req.method is invalid [req: %+v]", req)
		}

		if err != nil {
			c.lgr.Print(NewLog(
				LogLevelERROR,
				c.app.Hostname,
				serviceNameSimpleTwitterClient,
				err.Error(),
			))

			continue
		}

		// TODO
		fmt.Println(httpRes)
		c.resC <- nil
	}

	// Send a closed signal.
	c.closedReqC <- struct{}{}
}

// NewSimpleTwitterClient creates and returns a simple Twitter client.
func NewSimpleTwitterClient(
	consumerKey string,
	consumerSecret string,
	accessToken string,
	accessTokenSecret string,
	app *Application,
	lgr Logger,
	opts *SimpleTwitterClientOptions,
) TwitterClient {
	// Initialize options.
	if opts == nil {
		opts = &SimpleTwitterClientOptions{}
	}
	opts.setDefaults()

	c := &simpleTwitterClient{
		consumer: oauth.NewConsumer(
			consumerKey,
			consumerSecret,
			oauth.ServiceProvider{
				RequestTokenUrl:   twitterRequestTokenURL,
				AuthorizeTokenUrl: twitterAuthorizeTokenURL,
				AccessTokenUrl:    twitterAccessTokenURL,
			},
		),
		accessToken: &oauth.AccessToken{
			Token:          accessToken,
			Secret:         accessTokenSecret,
			AdditionalData: nil,
		},
		reqC:       make(chan twitterRequest, opts.ReqCBfSize),
		resC:       make(chan TwitterResponse, opts.ResCBfSize),
		closedReqC: make(chan struct{}),
		app:        app,
		lgr:        lgr,
	}

	// Launch a call goroutine.
	go c.call()

	return c
}
