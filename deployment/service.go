package deployment

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type DeploymentService struct {
	urlBase string
	client  *http.Client
}

func NewDeploymentService(url string, client *http.Client) *DeploymentService {
	return &DeploymentService{
		urlBase: url,
		client:  client,
	}
}

const URL = "http://localhost:3000"

func (ds *DeploymentService) Get(ID string) (*Deployment, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/deployments/%s", ds.urlBase, ID), nil)

	if err != nil {
		return nil, err
	}

	res, err := ds.client.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	dm := &Deployment{}
	err = json.Unmarshal(b, dm)

	if err != nil {
		return nil, err
	}

	return dm, err
}

func (ds *DeploymentService) Post(params DeploymentParams) (*Deployment, error) {
	body, err := json.Marshal(params)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/deployments", ds.urlBase), bytes.NewReader(body))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	res, err := ds.client.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	dm := &Deployment{}
	err = json.Unmarshal(b, dm)

	if err != nil {
		return nil, err
	}

	return dm, err
}

func (ds *DeploymentService) Delete(ID string) error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/deployments/%s", ds.urlBase, ID), nil)

	if err != nil {
		return err
	}

	_, err = ds.client.Do(req)

	if err != nil {
		return err
	}

	return nil
}
