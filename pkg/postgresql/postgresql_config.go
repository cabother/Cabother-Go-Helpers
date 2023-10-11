package postgresql

import (
	environment "github.com/cabother/cabother-go-helpers/pkg/so"
)

const (
	defaultHostName     = "POSTGRESQL_DB_HOST"
	defaultDatabaseName = "POSTGRESQL_DB_NAME"
	defaultUserName     = "POSTGRESQL_DB_USER"
	defaultPasswordName = "POSTGRESQL_DB_PASSWORD"
	defaultPort         = "POSTGRESQL_DB_PORT"

	defaultHostValue     = "127.0.0.1"
	defaultDatabaseValue = "cabothergohelpers"
	defaultUserValue     = "cabothergohelpers"
	defaultPasswordValue = "cabothergohelpers"
	defaultPortValue     = 5432
)

type Configuration interface {
	GetDefaultPostgreSQLConfigs() *PostgreSQLConfig
	GetCustomPostgreSQLConfigs(database string, host string, username string, password string) *PostgreSQLConfig
}

type configuration struct {
	env environment.Environment
}

func New(env environment.Environment) Configuration {
	return &configuration{env}
}

func (a *configuration) GetCustomPostgreSQLConfigs(database string,
	host string,
	username string,
	password string,
) *PostgreSQLConfig {
	return &PostgreSQLConfig{
		PostgreSQLDatabase: database,
		PostgreSQLHost:     host,
		PostgreSQLUsername: username,
		PostgreSQLPassword: password,
	}
}

func (a *configuration) GetDefaultPostgreSQLConfigs() *PostgreSQLConfig {
	database := a.env.GetEnvOrDefault(defaultDatabaseName, defaultDatabaseValue)
	host := a.env.GetEnvOrDefault(defaultHostName, defaultHostValue)
	user := a.env.GetEnvOrDefault(defaultUserName, defaultUserValue)
	password := a.env.GetEnvOrDefault(defaultPasswordName, defaultPasswordValue)
	port := a.env.GetIntEnvOrDefault(defaultPort, defaultPortValue)

	return &PostgreSQLConfig{
		PostgreSQLDatabase: database,
		PostgreSQLHost:     host,
		PostgreSQLUsername: user,
		PostgreSQLPassword: password,
		PostgreSQLPort:     port,
	}
}

type PostgreSQLConfig struct {
	PostgreSQLHost     string
	PostgreSQLDatabase string
	PostgreSQLUsername string
	PostgreSQLPassword string
	PostgreSQLPort     int
}
