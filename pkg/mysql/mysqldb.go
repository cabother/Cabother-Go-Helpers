package mysql

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const defaultMySQLPort = 3306

type MySQLService interface {
	Create(value interface{}) (tx *gorm.DB)
	CreateInBatches(value interface{}, batchSize int) (tx *gorm.DB)
	Save(value interface{}) (tx *gorm.DB)
	First(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	Take(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	Last(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	Find(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB
	FirstOrInit(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	FirstOrCreate(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	Update(column string, value interface{}) (tx *gorm.DB)
	Updates(values interface{}) (tx *gorm.DB)
	UpdateColumn(column string, value interface{}) (tx *gorm.DB)
	UpdateColumns(values interface{}) (tx *gorm.DB)
	Delete(value interface{}, conds ...interface{}) (tx *gorm.DB)
	Count(count *int64) (tx *gorm.DB)
	Row() *sql.Row
	Rows() (*sql.Rows, error)
	Scan(dest interface{}) (tx *gorm.DB)
	Pluck(column string, dest interface{}) (tx *gorm.DB)
	ScanRows(rows *sql.Rows, dest interface{}) error
	Connection(fc func(tx *gorm.DB) error) (err error)
	Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error)
	Begin(opts ...*sql.TxOptions) *gorm.DB
	Commit() *gorm.DB
	Rollback() *gorm.DB
	SavePoint(name string) *gorm.DB
	RollbackTo(name string) *gorm.DB
	Exec(sql string, values ...interface{}) (tx *gorm.DB)
}

type mySQLService struct {
	DB *gorm.DB
}

func NewMySQLService(config dbConfig) (MySQLService, error) {
	fmt.Println("[Cabother - Go Helpers] Connecting to MySql database...")

	dsn := config.toMySQLConnectionString()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &mySQLService{db}, nil
}

func NewMySQLServiceByDSN(dsn string) (MySQLService, error) {
	fmt.Println("[Cabother - Go Helpers] Connecting to MySql database by dsn...")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &mySQLService{db}, nil
}

func (a *mySQLService) Create(value interface{}) (tx *gorm.DB) {
	return a.DB.Create(value)
}

func (a *mySQLService) CreateInBatches(value interface{}, batchSize int) (tx *gorm.DB) {
	return a.DB.CreateInBatches(value, batchSize)
}

func (a *mySQLService) Save(value interface{}) (tx *gorm.DB) {
	return a.DB.Save(value)
}

func (a *mySQLService) First(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
	return a.DB.First(dest, conds...)
}

func (a *mySQLService) Take(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
	return a.DB.Take(dest, conds...)
}

func (a *mySQLService) Last(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
	return a.DB.Last(dest, conds...)
}

func (a *mySQLService) Find(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
	return a.DB.Find(dest, conds...)
}

func (a *mySQLService) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return a.DB.FindInBatches(dest, batchSize, fc)
}

func (a *mySQLService) FirstOrInit(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
	return a.DB.FirstOrInit(dest, conds...)
}

func (a *mySQLService) FirstOrCreate(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
	return a.DB.FirstOrCreate(dest, conds...)
}

func (a *mySQLService) Update(column string, value interface{}) (tx *gorm.DB) {
	return a.DB.Update(column, value)
}

func (a *mySQLService) Updates(values interface{}) (tx *gorm.DB) {
	return a.DB.Updates(values)
}

func (a *mySQLService) UpdateColumn(column string, value interface{}) (tx *gorm.DB) {
	return a.DB.UpdateColumn(column, value)
}

func (a *mySQLService) UpdateColumns(values interface{}) (tx *gorm.DB) {
	return a.DB.UpdateColumns(values)
}

func (a *mySQLService) Delete(value interface{}, conds ...interface{}) (tx *gorm.DB) {
	return a.DB.Delete(value, conds...)
}

func (a *mySQLService) Count(count *int64) (tx *gorm.DB) {
	return a.DB.Count(count)
}

func (a *mySQLService) Row() *sql.Row {
	return a.DB.Row()
}

func (a *mySQLService) Rows() (*sql.Rows, error) {
	return a.DB.Rows()
}

func (a *mySQLService) Scan(dest interface{}) (tx *gorm.DB) {
	return a.DB.Scan(dest)
}

func (a *mySQLService) Pluck(column string, dest interface{}) (tx *gorm.DB) {
	return a.DB.Pluck(column, dest)
}

func (a *mySQLService) ScanRows(rows *sql.Rows, dest interface{}) error {
	return a.DB.ScanRows(rows, dest)
}

func (a *mySQLService) Connection(fc func(tx *gorm.DB) error) (err error) {
	return a.DB.Connection(fc)
}

func (a *mySQLService) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return a.DB.Transaction(fc)
}

func (a *mySQLService) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return a.DB.Begin()
}

func (a *mySQLService) Commit() *gorm.DB {
	return a.DB.Commit()
}

func (a *mySQLService) Rollback() *gorm.DB {
	return a.DB.Rollback()
}

func (a *mySQLService) SavePoint(name string) *gorm.DB {
	return a.DB.SavePoint(name)
}

func (a *mySQLService) RollbackTo(name string) *gorm.DB {
	return a.DB.RollbackTo(name)
}

func (a *mySQLService) Exec(sql string, values ...interface{}) (tx *gorm.DB) {
	return a.DB.Exec(sql, values...)
}

type dbConfig struct {
	Port     int
	Host     string
	Database string
	Username string
	Password string
}

func (c *dbConfig) toMySQLConnectionString() string {
	if c.Port == 0 {
		c.Port = defaultMySQLPort
	}

	return fmt.Sprintf(
		"%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		c.Username,
		c.Password,
		c.Host,
		c.Port,
		c.Database,
	)
}
