package mysql

import (
	environment "github.com/cabother/cabother-go-helpers/pkg/so"
)

const (
	defaultDatabaseName = "MYSQL_DB_NAME"
	defaultHostName     = "MYSQL_DB_HOST"
	defaultUserName     = "MYSQL_DB_USER"
	defaultPasswordName = "MYSQL_DB_PASSWORD"

	defaultDatabaseValue = "cabothergohelpers"
	defaultHostValue     = "cabothergohelpers"
	defaultUserValue     = "cabothergohelpers"
	defaultPasswordValue = "cabothergohelpers"
)

type ConfigurationInterface interface {
	GetDefaultMySQLConfigs() *MySQLConfig
	GetCustomMySQLConfigs(database string, host string, username string, password string) *MySQLConfig
}

type Configuration struct {
	env environment.EnvironmentInterface
}

func NewConfiguration(env environment.EnvironmentInterface) ConfigurationInterface {
	return &Configuration{env}
}

func (a *Configuration) GetCustomMySQLConfigs(database string,
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

func (a *Configuration) GetDefaultMySQLConfigs() *MySQLConfig {
	database := a.env.GetEnvOrDefault(defaultDatabaseName, defaultDatabaseValue)
	host := a.env.GetEnvOrDefault(defaultHostName, defaultHostValue)
	user := a.env.GetEnvOrDefault(defaultUserName, defaultUserValue)
	password := a.env.GetEnvOrDefault(defaultPasswordName, defaultPasswordValue)

	return &MySQLConfig{
		MySQLDatabase: database,
		MySQLHost:     host,
		MySQLUsername: user,
		MySQLPassword: password,
	}
}

type MySQLConfig struct {
	MySQLHost     string
	MySQLDatabase string
	MySQLUsername string
	MySQLPassword string
}
