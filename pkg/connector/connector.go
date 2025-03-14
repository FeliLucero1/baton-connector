package connector

import (
	"context"
	"io"

	"github.com/conductorone/baton-debug-zone/pkg/client"
	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
	"github.com/conductorone/baton-sdk/pkg/annotations"
	"github.com/conductorone/baton-sdk/pkg/connectorbuilder"
)

type Connector struct {
	apiClient *client.APIClient
}

// ResourceSyncers returns a ResourceSyncer for each resource type that should be synced from the upstream service.
func (d *Connector) ResourceSyncers(ctx context.Context) []connectorbuilder.ResourceSyncer {
	apiClient := client.NewClient("your-username", "your-password") // Create an instance of APIClient
	return []connectorbuilder.ResourceSyncer{
		newUserBuilder(apiClient),
	}
}

// Asset takes an input AssetRef and attempts to fetch it using the connector's authenticated http client
// It streams a response, always starting with a metadata object, following by chunked payloads for the asset.
func (d *Connector) Asset(ctx context.Context, asset *v2.AssetRef) (string, io.ReadCloser, error) {
	return "", nil, nil
}

// Metadata returns metadata about the connector.
func (d *Connector) Metadata(ctx context.Context) (*v2.ConnectorMetadata, error) {
	return &v2.ConnectorMetadata{
		DisplayName: "My Baton Connector",
		Description: "The template implementation of a baton connector",
	}, nil
}

// Validate is called to ensure that the connector is properly configured. It should exercise any API credentials
// to be sure that they are valid.
func (d *Connector) Validate(ctx context.Context) (annotations.Annotations, error) {
	return nil, nil
}

// New crea una nueva instancia del conector con autenticación
func New(ctx context.Context, apiBaseURL, username, password string) (*Connector, error) {
	apiClient := client.NewClient(username, password)
	err := apiClient.SetBaseURL(apiBaseURL)
	if err != nil {
		return nil, err
	}

	return &Connector{apiClient: apiClient}, nil
}
