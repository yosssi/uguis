package uguis

// HTTP methods
const (
	httpGET    = "GET"
	httpPOST   = "POST"
	httpDELETE = "DELETE"
)

// twitterRequest represents a request for a Twitter API.
type twitterRequest struct {
	// method represents an HTTP request.
	method string
	// url represents an URL.
	path string
	// params represents request parameters.
	params map[string]string
}

// newTwitterRequest creates and returns a request for a Twitter API.
func newTwitterRequest(method string, path string, params map[string]string) twitterRequest {
	return twitterRequest{
		method: method,
		path:   path,
		params: params,
	}
}
