package postgresql

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const defaultPostgreSQLPort = 5432

type PostgreSQLService interface {
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

type postgreSQLService struct {
	DB *gorm.DB
}

func NewPostgreSQLService(config dbConfig) (PostgreSQLService, error) {
	fmt.Println("[Cabother - Go Helpers] Connecting to PostgreSql database...")

	dsn := config.toPostgreSQLConnectionString()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &postgreSQLService{db}, nil
}

func NewPostgreSQLServiceByDSN(dsn string) (PostgreSQLService, error) {
	fmt.Println("[Cabother - Go Helpers] Connecting to postgre database by dsn...")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &postgreSQLService{db}, nil
}

func (a *postgreSQLService) Create(value interface{}) (tx *gorm.DB) {
	return a.DB.Create(value)
}

func (a *postgreSQLService) CreateInBatches(value interface{}, batchSize int) (tx *gorm.DB) {
	return a.DB.CreateInBatches(value, batchSize)
}

func (a *postgreSQLService) Save(value interface{}) (tx *gorm.DB) {
	return a.DB.Save(value)
}

func (a *postgreSQLService) First(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
	return a.DB.First(dest, conds...)
}

func (a *postgreSQLService) Take(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
	return a.DB.Take(dest, conds...)
}

func (a *postgreSQLService) Last(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
	return a.DB.Last(dest, conds...)
}

func (a *postgreSQLService) Find(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
	return a.DB.Find(dest, conds...)
}

func (a *postgreSQLService) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return a.DB.FindInBatches(dest, batchSize, fc)
}

func (a *postgreSQLService) FirstOrInit(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
	return a.DB.FirstOrInit(dest, conds...)
}

func (a *postgreSQLService) FirstOrCreate(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
	return a.DB.FirstOrCreate(dest, conds...)
}

func (a *postgreSQLService) Update(column string, value interface{}) (tx *gorm.DB) {
	return a.DB.Update(column, value)
}

func (a *postgreSQLService) Updates(values interface{}) (tx *gorm.DB) {
	return a.DB.Updates(values)
}

func (a *postgreSQLService) UpdateColumn(column string, value interface{}) (tx *gorm.DB) {
	return a.DB.UpdateColumn(column, value)
}

func (a *postgreSQLService) UpdateColumns(values interface{}) (tx *gorm.DB) {
	return a.DB.UpdateColumns(values)
}

func (a *postgreSQLService) Delete(value interface{}, conds ...interface{}) (tx *gorm.DB) {
	return a.DB.Delete(value, conds...)
}

func (a *postgreSQLService) Count(count *int64) (tx *gorm.DB) {
	return a.DB.Count(count)
}

func (a *postgreSQLService) Row() *sql.Row {
	return a.DB.Row()
}

func (a *postgreSQLService) Rows() (*sql.Rows, error) {
	return a.DB.Rows()
}

func (a *postgreSQLService) Scan(dest interface{}) (tx *gorm.DB) {
	return a.DB.Scan(dest)
}

func (a *postgreSQLService) Pluck(column string, dest interface{}) (tx *gorm.DB) {
	return a.DB.Pluck(column, dest)
}

func (a *postgreSQLService) ScanRows(rows *sql.Rows, dest interface{}) error {
	return a.DB.ScanRows(rows, dest)
}

func (a *postgreSQLService) Connection(fc func(tx *gorm.DB) error) (err error) {
	return a.DB.Connection(fc)
}

func (a *postgreSQLService) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return a.DB.Transaction(fc)
}

func (a *postgreSQLService) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return a.DB.Begin()
}

func (a *postgreSQLService) Commit() *gorm.DB {
	return a.DB.Commit()
}

func (a *postgreSQLService) Rollback() *gorm.DB {
	return a.DB.Rollback()
}

func (a *postgreSQLService) SavePoint(name string) *gorm.DB {
	return a.DB.SavePoint(name)
}

func (a *postgreSQLService) RollbackTo(name string) *gorm.DB {
	return a.DB.RollbackTo(name)
}

func (a *postgreSQLService) Exec(sql string, values ...interface{}) (tx *gorm.DB) {
	return a.DB.Exec(sql, values...)
}

type dbConfig struct {
	Port     int
	Host     string
	Database string
	Username string
	Password string
}

func (c *dbConfig) toPostgreSQLConnectionString() string {
	if c.Port == 0 {
		c.Port = defaultPostgreSQLPort
	}

	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%v sslmode=disable",
		c.Host,
		c.Username,
		c.Password,
		c.Database,
		c.Port,
	)
}
