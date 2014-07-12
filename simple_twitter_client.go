package uguis

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mrjones/oauth"
)

// Twitter API URLs
const (
	twitterRequestTokenURL   = "http://api.twitter.com/oauth/request_token"
	twitterAuthorizeTokenURL = "https://api.twitter.com/oauth/authorize"
	twitterAccessTokenURL    = "https://api.twitter.com/oauth/access_token"
	twitterBaseURL           = "https://api.twitter.com/1.1"
)

const serviceNameSimpleTwitterClient = "simpleTwitterClient"

// simpleTwitterClient represents a simple Twitter client.
type simpleTwitterClient struct {
	consumer    *oauth.Consumer
	accessToken *oauth.AccessToken
	reqC        chan twitterRequest
	resC        chan Tweet
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

		url := twitterBaseURL + req.path

		switch req.method {
		case httpGET:
			httpRes, err = c.consumer.Get(url, req.params, c.accessToken)
		case httpPOST:
			httpRes, err = c.consumer.Post(url, req.params, c.accessToken)
		case httpDELETE:
			httpRes, err = c.consumer.Delete(url, req.params, c.accessToken)
		default:
			err = fmt.Errorf("req.method is invalid [req: %+v]", req)
		}

		if err != nil {
			c.logError(err)
			continue
		}

		// Parse the response.
		b, err := ioutil.ReadAll(httpRes.Body)
		httpRes.Body.Close()
		if err != nil {
			c.logError(err)
			continue
		}

		var tweets Tweets
		if err := json.Unmarshal(b, &tweets); err != nil {
			c.logError(err)
			continue
		}

		for i := len(tweets) - 1; i >= 0; i-- {
			c.resC <- tweets[i]
		}
	}

	// Send a closed signal.
	c.closedReqC <- struct{}{}
}

// ResC returns a response channel.
func (c *simpleTwitterClient) ResC() <-chan Tweet {
	return c.resC
}

func (c *simpleTwitterClient) logError(err error) {
	c.lgr.Print(NewLog(
		LogLevelERROR,
		c.app.Hostname,
		serviceNameSimpleTwitterClient,
		err.Error(),
	))
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
		resC:       make(chan Tweet, opts.ResCBfSize),
		closedReqC: make(chan struct{}),
		app:        app,
		lgr:        lgr,
	}

	// Launch a call goroutine.
	go c.call()

	return c
}
