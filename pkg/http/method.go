package http

import (
	"bytes"
	"net/http"
)

func (c *CustomHTTPClient) Get(path string) (*http.Response, error) {
	fullURL := c.BaseURL + path

	return c.Client.Get(fullURL)
}

func (c *CustomHTTPClient) Post(path string, contentType string, body []byte) (*http.Response, error) {
	fullURL := c.BaseURL + path

	return c.Client.Post(fullURL, contentType, bytes.NewBuffer(body))
}

func (c *CustomHTTPClient) Patch(path string, contentType string, body []byte) (*http.Response, error) {
	fullURL := c.BaseURL + path

	req, err := http.NewRequest(http.MethodPatch, fullURL, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", contentType)

	if transport, ok := c.Client.Transport.(*defaultTransport); ok {
		for key, value := range transport.Headers {
			req.Header.Set(key, value)
		}
	}

	if transport, ok := c.Client.Transport.(*defaultTransport); ok && transport.Authorization != "" {
		req.Header.Set("Authorization", transport.Authorization)
	}

	return c.Client.Do(req)
}
