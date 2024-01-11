package http

import (
	"net/http"
	"time"

	"github.com/esc-chula/esc-docsync/pkg/config"
)

type CustomHTTPClient struct {
	Client  *http.Client
	BaseURL string
}

func NewCustomHTTPClient() *CustomHTTPClient {
	appConfig := config.AppConfig()

	transport := &defaultTransport{
		Transport: http.DefaultTransport,
	}

	return &CustomHTTPClient{
		Client: &http.Client{
			Transport: transport,
			Timeout:   time.Second * time.Duration(appConfig.RequestTimeout),
		},
		BaseURL: "",
	}
}

func (c *CustomHTTPClient) SetBaseURL(baseURL string) {
	c.BaseURL = baseURL
}

func (c *CustomHTTPClient) SetDefaultHeaders(headers map[string]string) {
	transport, ok := c.Client.Transport.(*defaultTransport)
	if !ok {
		transport = &defaultTransport{
			Transport: c.Client.Transport,
		}
		c.Client.Transport = transport
	}

	transport.Headers = headers
}

func (c *CustomHTTPClient) SetBearerAuthorization(token string) {
	transport, ok := c.Client.Transport.(*defaultTransport)
	if !ok {
		transport = &defaultTransport{
			Transport: c.Client.Transport,
		}
		c.Client.Transport = transport
	}

	transport.Authorization = "Bearer " + token
}

type defaultTransport struct {
	Transport     http.RoundTripper
	BaseURL       string
	Headers       map[string]string
	Authorization string
}

func (dt *defaultTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	for key, value := range dt.Headers {
		req.Header.Set(key, value)
	}

	if dt.Authorization != "" {
		req.Header.Set("Authorization", dt.Authorization)
	}

	if dt.BaseURL != "" {
		req.URL.Path = dt.BaseURL + req.URL.Path
	}

	return dt.Transport.RoundTrip(req)
}
