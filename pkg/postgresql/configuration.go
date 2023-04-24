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

	defaultHostValue     = "127.0.0.1ððð"
	defaultDatabaseValue = "cabothergohelpers"
	defaultUserValue     = "cabothergohelpers"
	defaultPasswordValue = "cabothergohelpers"
	defaultPortValue     = 5432
)

type ConfigurationInterface interface {
	GetDefaultPostgreSQLConfigs() *MySQLConfig
	GetCustomPostgreSQLConfigs(database string, host string, username string, password string) *MySQLConfig
}

type Configuration struct {
	env environment.EnvironmentInterface
}

func New(env environment.EnvironmentInterface) ConfigurationInterface {
	return &Configuration{env}
}

func (a *Configuration) GetCustomPostgreSQLConfigs(database string,
	host string,
	username string,
	password string,
) *MySQLConfig {
	return &MySQLConfig{
		PostgreSQLDatabase: database,
		PostgreSQLHost:     host,
		PostgreSQLUsername: username,
		PostgreSQLPassword: password,
	}
}

func (a *Configuration) GetDefaultPostgreSQLConfigs() *MySQLConfig {
	database := a.env.GetEnvOrDefault(defaultDatabaseName, defaultDatabaseValue)
	host := a.env.GetEnvOrDefault(defaultHostName, defaultHostValue)
	user := a.env.GetEnvOrDefault(defaultUserName, defaultUserValue)
	password := a.env.GetEnvOrDefault(defaultPasswordName, defaultPasswordValue)
	port := a.env.GetIntEnvOrDefault(defaultPort, defaultPortValue)

	return &MySQLConfig{
		PostgreSQLDatabase: database,
		PostgreSQLHost:     host,
		PostgreSQLUsername: user,
		PostgreSQLPassword: password,
		PostgreSQLPort:     port,
	}
}

type MySQLConfig struct {
	PostgreSQLHost     string
	PostgreSQLDatabase string
	PostgreSQLUsername string
	PostgreSQLPassword string
	PostgreSQLPort     int
}
