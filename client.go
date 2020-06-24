package webfinger

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

// DefaultPath to webfinger service
const DefaultPath = "/.well-known/webfinger"

// Client is used to query a webfinger server
type Client struct {
	cli *http.Client
	url url.URL
}

// ClientOption client options func
type ClientOption func(c *Client)

// WithHTTPClient sets the http client of the client to other
// than http.DefaultClient
func WithHTTPClient(cli *http.Client) ClientOption {
	return func(c *Client) {
		c.cli = cli
	}
}

// NewClient returns a new webfinger client instance
func NewClient(u string, opts ...ClientOption) (*Client, error) {

	uu, err := url.Parse(u)
	if err != nil {
		return nil, err
	}

	c := &Client{
		cli: http.DefaultClient,
		url: *uu,
	}

	for _, o := range opts {
		o(c)
	}

	return c, nil
}

// Query the server for a resource
func (c *Client) Query(q *Query) (*Resource, error) {

	u := c.url
	u.RawQuery = q.ToValues().Encode()

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	d, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	r := &Resource{}

	return r, json.Unmarshal(d, r)
}
