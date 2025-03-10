package client

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/conductorone/baton-sdk/pkg/uhttp"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
)

const (
	baseDomain        = "https://your-api.com/api"
	getUsersEndpoint  = "/users"
	getUserByID       = "/users/%d"
	getProjects       = "/projects"
	getProjectByID    = "/projects/%d"
	getUsersByProject = "/projects/%d/users"
	getRoles          = "/roles"
)

type APIClient struct {
	//basic auth
	username string
	password string
	//url base de la API
	baseURL string
	//wrapper para hacer las peticiones HTTP
	wrapper *uhttp.BaseHttpClient
}

// creacion del cliente
func NewClient(username, password string, httpClient ...*uhttp.BaseHttpClient) *APIClient {
	var wrapper = &uhttp.BaseHttpClient{}
	if len(httpClient) > 0 {
		wrapper = httpClient[0]
	}
	return &APIClient{
		wrapper:  wrapper,
		baseURL:  baseDomain,
		username: username,
		password: password,
	}
}

func (c *APIClient) ListUsers(ctx context.Context) ([]User, error) {
	var res []User
	queryUrl, err := url.JoinPath(c.baseURL, getUsersEndpoint)
	if err != nil {
		return nil, err
	}

	err = c.getResourcesFromAPI(ctx, queryUrl, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *APIClient) GetUserByID(ctx context.Context, id int) (*User, error) {
	queryUrl := fmt.Sprintf(c.baseURL+getUserByID, id)
	var res *User
	err := c.getResourcesFromAPI(ctx, queryUrl, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *APIClient) getResourcesFromAPI(ctx context.Context, urlAddress string, res any) error {
	_, err := c.doRequest(ctx, http.MethodGet, urlAddress, &res)
	return err
}

// Función para hacer peticiones
func (c *APIClient) doRequest(ctx context.Context, method string, endpointUrl string, res interface{}) (http.Header, error) {
	l := ctxzap.Extract(ctx)

	req, err := http.NewRequestWithContext(ctx, method, endpointUrl, nil)
	if err != nil {
		l.Error(fmt.Sprintf("Error creating request: %s", err))
		return nil, err
	}

	req.Header.Set("Authorization", c.basicAuthHeader())
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := c.wrapper.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if res != nil {
		err = json.NewDecoder(resp.Body).Decode(&res)
		if err != nil {
			return nil, err
		}
	}

	return resp.Header, nil
}

func (c *APIClient) basicAuthHeader() string {
	auth := c.username + ":" + c.password
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
}

func (c *APIClient) ListProjects(ctx context.Context) ([]Project, error) {
	var res []Project
	queryUrl, err := url.JoinPath(c.baseURL, getProjects)
	if err != nil {
		return nil, err
	}

	err = c.getResourcesFromAPI(ctx, queryUrl, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *APIClient) GetProjectByID(ctx context.Context, id int) (*Project, error) {
	queryUrl := fmt.Sprintf(c.baseURL+getProjectByID, id)
	var res *Project
	err := c.getResourcesFromAPI(ctx, queryUrl, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *APIClient) ListUsersByProject(ctx context.Context, projectID int) ([]UserWithRole, error) {
	var res []UserWithRole
	queryUrl := fmt.Sprintf(c.baseURL+getUsersByProject, projectID)
	err := c.getResourcesFromAPI(ctx, queryUrl, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *APIClient) ListRoles(ctx context.Context) ([]string, error) {
	var res []string
	queryUrl, err := url.JoinPath(c.baseURL, getRoles)
	if err != nil {
		return nil, err
	}

	err = c.getResourcesFromAPI(ctx, queryUrl, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *APIClient) SetBaseURL(newBaseURL string) error {
	parsedURL, err := url.ParseRequestURI(newBaseURL)
	if err != nil {
		return fmt.Errorf("invalid base URL: %w", err)
	}
	c.baseURL = parsedURL.String()
	return nil
}
