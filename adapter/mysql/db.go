package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/originbenntou/modev-backend/config"
	"github.com/pkg/errors"
	"time"
)

type DB struct {
	*sqlx.DB
	Name string
}

func NewDB() (*DB, error) {
	c, err := connect(connectString(config.GetConfig().Database))
	if err != nil {
		return nil, errors.Wrap(err, "db connect failed")
	}

	return &DB{
		DB:   c,
		Name: config.DefaultDBName,
	}, nil
}

func connect(connStr string) (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}

	setupDBConns(db)

	return db, nil
}

func setupDBConns(db *sqlx.DB) {
	db.SetConnMaxLifetime(config.ConnMaxLifetimeSec * time.Second)
	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetMaxIdleConns(config.MaxIdleConns)
}

func connectString(database config.Database) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=true",
		database.User,
		database.Password,
		database.Host,
		database.Port,
		database.Name,
	)
}
