package mysql

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLServiceInterface interface {
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

func (a *MySQLService) Create(value interface{}) (tx *gorm.DB) {
	return a.DB.Create(value)
}

func (a *MySQLService) CreateInBatches(value interface{}, batchSize int) (tx *gorm.DB) {
	return a.DB.CreateInBatches(value, batchSize)
}

func (a *MySQLService) Save(value interface{}) (tx *gorm.DB) {
	return a.DB.Save(value)
}

func (a *MySQLService) First(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
	return a.DB.First(dest, conds...)
}

func (a *MySQLService) Take(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
	return a.DB.Take(dest, conds...)
}

func (a *MySQLService) Last(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
	return a.DB.Last(dest, conds...)
}

func (a *MySQLService) Find(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
	return a.DB.Find(dest, conds...)
}

func (a *MySQLService) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return a.DB.FindInBatches(dest, batchSize, fc)
}

func (a *MySQLService) FirstOrInit(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
	return a.DB.FirstOrInit(dest, conds...)
}

func (a *MySQLService) FirstOrCreate(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
	return a.DB.FirstOrCreate(dest, conds...)
}

func (a *MySQLService) Update(column string, value interface{}) (tx *gorm.DB) {
	return a.DB.Update(column, value)
}

func (a *MySQLService) Updates(values interface{}) (tx *gorm.DB) {
	return a.DB.Updates(values)
}

func (a *MySQLService) UpdateColumn(column string, value interface{}) (tx *gorm.DB) {
	return a.DB.UpdateColumn(column, value)
}

func (a *MySQLService) UpdateColumns(values interface{}) (tx *gorm.DB) {
	return a.DB.UpdateColumns(values)
}

func (a *MySQLService) Delete(value interface{}, conds ...interface{}) (tx *gorm.DB) {
	return a.DB.Delete(value, conds...)
}

func (a *MySQLService) Count(count *int64) (tx *gorm.DB) {
	return a.DB.Count(count)
}

func (a *MySQLService) Row() *sql.Row {
	return a.DB.Row()
}

func (a *MySQLService) Rows() (*sql.Rows, error) {
	return a.DB.Rows()
}

func (a *MySQLService) Scan(dest interface{}) (tx *gorm.DB) {
	return a.DB.Scan(dest)
}

func (a *MySQLService) Pluck(column string, dest interface{}) (tx *gorm.DB) {
	return a.DB.Pluck(column, dest)
}

func (a *MySQLService) ScanRows(rows *sql.Rows, dest interface{}) error {
	return a.DB.ScanRows(rows, dest)
}

func (a *MySQLService) Connection(fc func(tx *gorm.DB) error) (err error) {
	return a.DB.Connection(fc)
}

func (a *MySQLService) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return a.DB.Transaction(fc)
}

func (a *MySQLService) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return a.DB.Begin()
}

func (a *MySQLService) Commit() *gorm.DB {
	return a.DB.Commit()
}

func (a *MySQLService) Rollback() *gorm.DB {
	return a.DB.Rollback()
}

func (a *MySQLService) SavePoint(name string) *gorm.DB {
	return a.DB.SavePoint(name)
}

func (a *MySQLService) RollbackTo(name string) *gorm.DB {
	return a.DB.RollbackTo(name)
}

func (a *MySQLService) Exec(sql string, values ...interface{}) (tx *gorm.DB) {
	return a.DB.Exec(sql, values...)
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
