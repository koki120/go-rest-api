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

const commonColumns = `
	created_at TIMESTAMP DEFAULT current_timestamp,
	updated_at TIMESTAMP DEFAULT current_timestamp ON UPDATE CURRENT_TIMESTAMP
`

func Migrate(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			hashed_password VARCHAR(255) NOT NULL,
			user_type ENUM('SystemAdmin', 'User') DEFAULT 'User',
			` + commonColumns + `
		);
	`)
	if err != nil {
		return err
	}
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS memos (
			id SERIAL PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			body TEXT,
			user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
			` + commonColumns + `
		);
	`)
	if err != nil {
		return err
	}
	return nil
}

func dropAllTables(db *sql.DB) error {
	rows, err := db.Query("SHOW TABLES")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			return err
		}

		query := fmt.Sprintf("DROP TABLE IF EXISTS %s", tableName)
		if _, err := db.Exec(query); err != nil {
			return err
		}
	}

	return nil
}
