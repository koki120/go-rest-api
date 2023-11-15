package database

import (
	"database/sql"
	"fmt"
	"math"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/koki120/go-rest-api/config"
	"github.com/koki120/go-rest-api/log"
)

func NewMySQLDB() (*sql.DB, error) {
	logger := log.NewLogger()

	db, err := sql.Open("mysql", config.DSN())
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetMaxOpenConns(10)

	var communicationAttempts int = 5
	for i := 0; i < communicationAttempts; i++ {
		err = db.Ping()
		if err == nil {
			break
		}
		if i == (communicationAttempts - 1) {
			return nil, fmt.Errorf("failed to connect to database: %w", err)
		}
		duration := time.Millisecond * time.Duration(math.Pow(1.5, float64(i))*1000)
		logger.Warn("failed to connect to database retrying")
		time.Sleep(duration)
	}

	return db, nil
}
