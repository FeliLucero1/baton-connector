package main

import (
	"context"
	"fmt"
	"os"

	"github.com/conductorone/baton-debug-zone/pkg/connector"
	"github.com/conductorone/baton-sdk/pkg/config"
	"github.com/conductorone/baton-sdk/pkg/connectorbuilder"
	"github.com/conductorone/baton-sdk/pkg/field"
	"github.com/conductorone/baton-sdk/pkg/types"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var version = "dev"

func main() {
	ctx := context.Background()

	_, cmd, err := config.DefineConfiguration(
		ctx,
		"baton-debug-zone",
		getConnector,
		field.Configuration{
			Fields: ConfigurationFields,
		},
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	cmd.Version = version

	err = cmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func getConnector(ctx context.Context, v *viper.Viper) (types.ConnectorServer, error) {
	l := ctxzap.Extract(ctx)

	apiBaseURL := v.GetString(apiBaseURLField.FieldName)
	username := v.GetString(usernameField.FieldName)
	password := v.GetString(passwordField.FieldName)

	if err := ValidateConfig(v); err != nil {
		return nil, err
	}

	fmt.Println("API Base URL:", apiBaseURL)
	fmt.Println("Username:", username)
	fmt.Println("Password:", password)

	// Crear instancia del conector con los datos de autenticación
	connectorBuilder, err := connector.New(ctx, apiBaseURL, username, password)
	if err != nil {
		l.Error("error creating connector", zap.Error(err))
		return nil, err
	}

	connector, err := connectorbuilder.NewConnector(ctx, connectorBuilder)
	if err != nil {
		l.Error("error creating connector", zap.Error(err))
		return nil, err
	}
	return connector, nil
}
