package mysql

import (
	environment "github.com/cabother/cabother-go-helpers/pkg/so"
)

const (
	defaultHostName     = "MYSQL_DB_HOST"
	defaultDatabaseName = "MYSQL_DB_NAME"
	defaultUserName     = "MYSQL_DB_USER"
	defaultPasswordName = "MYSQL_DB_PASSWORD"
	defaultPort         = "MYSQL_DB_PORT"

	defaultHostValue     = "127.0.0.1"
	defaultDatabaseValue = "cabothergohelpers"
	defaultUserValue     = "cabothergohelpers"
	defaultPasswordValue = "cabothergohelpers"
	defaultPortValue     = 3306
)

type Configuration interface {
	GetDefaultMySQLConfigs() *MySQLConfig
	GetCustomMySQLConfigs(database string, host string, username string, password string) *MySQLConfig
}

type configuration struct {
	env environment.Environment
}

func New(env environment.Environment) Configuration {
	return &configuration{env}
}

func (a *configuration) GetCustomMySQLConfigs(database string,
	host string,
	username string,
	password string,
) *MySQLConfig {
	return &MySQLConfig{
		MySQLDatabase: database,
		MySQLHost:     host,
		MySQLUsername: username,
		MySQLPassword: password,
	}
}

func (a *configuration) GetDefaultMySQLConfigs() *MySQLConfig {
	database := a.env.GetEnvOrDefault(defaultDatabaseName, defaultDatabaseValue)
	host := a.env.GetEnvOrDefault(defaultHostName, defaultHostValue)
	user := a.env.GetEnvOrDefault(defaultUserName, defaultUserValue)
	password := a.env.GetEnvOrDefault(defaultPasswordName, defaultPasswordValue)
	port := a.env.GetIntEnvOrDefault(defaultPort, defaultPortValue)

	return &MySQLConfig{
		MySQLDatabase: database,
		MySQLHost:     host,
		MySQLUsername: user,
		MySQLPassword: password,
		MySQLPort:     port,
	}
}

type MySQLConfig struct {
	MySQLHost     string
	MySQLDatabase string
	MySQLUsername string
	MySQLPassword string
	MySQLPort     int
}
