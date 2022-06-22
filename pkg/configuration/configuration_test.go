package configuration

import (
	"os"
	"testing"

	environment "github.com/cabother/cabother-go-helpers/pkg/so"
	"github.com/stretchr/testify/assert"
)

const (
	mysqlStringDBPort = "3306"
	mysqlIntDBPort    = 3306
)

func setEnvs() {
	os.Setenv("MYSQL_DB_NAME", defaultDatabaseValue)
	os.Setenv("MYSQL_DB_HOST", defaultHostValue)
	os.Setenv("MYSQL_DB_USER", defaultUserValue)
	os.Setenv("MYSQL_DB_PASSWORD", defaultPasswordValue)
	os.Setenv("MYSQL_DB_PORT", mysqlStringDBPort)
}

func unsetEnvs() {
	os.Setenv("MYSQL_DB_NAME", "")
	os.Setenv("MYSQL_DB_HOST", "")
	os.Setenv("MYSQL_DB_USER", "")
	os.Setenv("MYSQL_DB_PASSWORD", "")
	os.Setenv("MYSQL_DB_PORT", "")
}

func Test_GetDefaultMySQLConfigs_Success(t *testing.T) {
	setEnvs()

	mysqlConfigs := NewConfiguration(environment.NewEnvironment()).GetDefaultMySQLConfigs()

	unsetEnvs()

	assert.Equal(t, defaultDatabaseValue, mysqlConfigs.MySQLDatabase)
	assert.Equal(t, defaultHostValue, mysqlConfigs.MySQLHost)
	assert.Equal(t, defaultUserValue, mysqlConfigs.MySQLUsername)
	assert.Equal(t, defaultPasswordValue, mysqlConfigs.MySQLPassword)
	assert.Equal(t, mysqlIntDBPort, mysqlConfigs.MySQLPort)
}
