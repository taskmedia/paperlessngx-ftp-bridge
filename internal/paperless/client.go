// Package paperless provides a minimal client for uploading documents to the
// paperless-ngx "post_document" API.
package paperless

import (
	"bytes"
	"fmt"

	"github.com/go-resty/resty/v2"
)

// Client uploads documents to a paperless-ngx instance via its REST API.
type Client struct {
	baseURL  string
	apiURL   string
	username string
	password string
	http     *resty.Client
}

// NewClient builds a Client that posts documents to
// "<baseURL>/api/documents/post_document/" using HTTP basic auth.
func NewClient(baseURL, username, password string) *Client {
	return &Client{
		baseURL:  baseURL,
		apiURL:   baseURL + "/api/documents/post_document/",
		username: username,
		password: password,
		http:     resty.New(),
	}
}

// Upload posts data to paperless-ngx as a document named filename. It
// returns an error if the request fails or paperless-ngx responds with a
// non-2xx status.
func (c *Client) Upload(filename string, data []byte) error {
	resp, err := c.http.R().
		SetBasicAuth(c.username, c.password).
		SetFileReader("document", filename, bytes.NewReader(data)).
		Post(c.apiURL)
	if err != nil {
		return fmt.Errorf("uploading document to paperless-ngx: %w", err)
	}

	if resp.IsError() {
		return fmt.Errorf("paperless-ngx rejected document %q: %s", filename, resp.Status())
	}

	return nil
}

// Ping checks whether paperless-ngx is reachable. Any HTTP response -- even
// a non-2xx one, e.g. an auth failure -- counts as reachable: the point is
// to detect a network-level outage (DNS failure, connection refused,
// timeout), not to validate credentials.
func (c *Client) Ping() error {
	if _, err := c.http.R().Get(c.baseURL); err != nil {
		return fmt.Errorf("checking paperless-ngx reachability: %w", err)
	}

	return nil
}
