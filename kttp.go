package kttp

import (
	"net/http"
	"time"

	"github.com/deborafelix/client-http/deployment"
)

type Client struct {
	url        string
	httpClient *http.Client
	timeout    time.Duration

	Deployment *deployment.DeploymentService
}

func NewClient(url string, timeout time.Duration) *Client {
	httpClient := &http.Client{
		Timeout: timeout,
	}
	return &Client{
		url:        url,
		timeout:    timeout,
		httpClient: httpClient,
		Deployment: deployment.NewDeploymentService(url, httpClient),
	}
}
