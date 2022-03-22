package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLServiceInterface interface{}

type MySQLService struct {
	DB *gorm.DB
}

func NewMySQLService(config DBConfig) (MySQLServiceInterface, error) {
	fmt.Println("Connecting to MySql database...")

	dsn := config.toMySQLConnectionString()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &MySQLService{db}, nil
}

type DBConfig struct {
	Host     string
	Database string
	Username string
	Password string
}

func (c *DBConfig) toMySQLConnectionString() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		c.Username,
		c.Password,
		c.Host,
		c.Database,
	)
}
