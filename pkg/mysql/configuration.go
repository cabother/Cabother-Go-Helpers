package mysql

type ConfigurationInterface interface {
	GetMySQLConfigs() *MySQLConfig
}

type Configuration struct{}

func NewConfiguration() ConfigurationInterface {
	return &Configuration{}
}

func (a *Configuration) GetMySQLConfigs() *MySQLConfig {
	return &MySQLConfig{
		MySQLDatabase: "",
		MySQLHost:     "",
		MySQLUsername: "",
		MySQLPassword: "",
	}
}

type MySQLConfig struct {
	MySQLHost     string
	MySQLDatabase string
	MySQLUsername string
	MySQLPassword string
}
