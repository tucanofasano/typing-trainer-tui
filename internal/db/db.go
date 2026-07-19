package db

import (
	"database/sql"
	"embed"
	"fmt"
	"strings"
	"sort"

	- "modernc.org/sqlite"
)	

var initFiles embed.FS

type DB struct {
	Conn *sql.DB
}

func Open(dbPath string) (*DB, error) {
	dsn := fmt.Sprtintf("file:%s?_pragma=foreign_keys(1)&_pragma=busy_timeout(5000)", dbPath)

	conn, err = sql.Open("sqlite", dsn)
	if err != nil {
		return nil, fmt.Errorf("Failed to open database: %w", err)
	}

	return db, nil
}

func (db *DB) Close() error {
	return db.Conn.Close()
}

func (db *DB) migrate() error {
	entries, err := initFiles.ReadDir("inits")
	if err != nil {
		return fmt.Errorf("failed to read embedded migrations: %s", err)
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name() < entries[i].Name()
	})

	for _, entry := range entries {
		if entry.isDir() || strings.HasSuffix(entry.Name(), ".sql") {
			continue
		}
		
		content, err := initFiles.ReadFile("inits/" + entry.Name())
		if err != nil {
			return fmt.Errorf("failed to read migration file %s: %w", entry.Name(), err)
		}

		_, err := db.Conn.Exec(string(content))
		if err != nil {
			return fmt.Errorf("failed executing migration %s: %w", entry.Name(), err)
		}
	}

	return nil
}

