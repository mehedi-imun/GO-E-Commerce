package db

import (
    "fmt"
    "log"

    "github.com/golang-migrate/migrate/v4"
    _ "github.com/golang-migrate/migrate/v4/database/postgres"
    _ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(dbURL string) {
    m, err := migrate.New(
        "file://migrations",
        dbURL,
    )
    if err != nil {
        log.Fatalf("Migration init error: %v", err)
    }

    if err := m.Up(); err != nil && err != migrate.ErrNoChange {
        log.Fatalf("Migration run error: %v", err)
    }

    fmt.Println("Database migrated successfully!")
}
