package main

import (
	"github.com/conductorone/baton-sdk/pkg/field"
	"github.com/spf13/viper"
)

var (
	apiBaseURLField = field.StringField(
		"api_base_url",
		field.WithDescription("Base URL de la API"),
		field.WithRequired(true),
	)

	usernameField = field.StringField(
		"username",
		field.WithDescription("Usuario para autenticación"),
		field.WithRequired(true),
	)

	passwordField = field.StringField(
		"password",
		field.WithDescription("Contraseña para autenticación"),
		field.WithIsSecret(true),
		field.WithRequired(true),
	)

	// Lista de campos de configuración requeridos
	ConfigurationFields = []field.SchemaField{apiBaseURLField, usernameField, passwordField}

	// No hay relaciones entre los campos en este caso
	FieldRelationships = []field.SchemaFieldRelationship{}

	cfg = field.Configuration{
		Fields:      ConfigurationFields,
		Constraints: FieldRelationships,
	}
)

// En trello no hicieron validación de la configuración
func ValidateConfig(v *viper.Viper) error {
	return nil
}
