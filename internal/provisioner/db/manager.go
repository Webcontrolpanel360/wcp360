package db

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

// CreateDatabase handles the creation of a new DB and User for a tenant
func CreateDatabase(dbName, username, password string) error {
    // Note: In production, use root credentials from a config file
    dsn := "root:your_root_password@tcp(127.0.0.1:3306)/"
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return err
    }
    defer db.Close()

    // SQL Commands
    queries := []string{
        fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName),
        fmt.Sprintf("CREATE USER IF NOT EXISTS '%s'@'localhost' IDENTIFIED BY '%s'", username, password),
        fmt.Sprintf("GRANT ALL PRIVILEGES ON %s.* TO '%s'@'localhost'", dbName, username),
        "FLUSH PRIVILEGES",
    }

    for _, q := range queries {
        _, err := db.Exec(q)
        if err != nil {
            return fmt.Errorf("error executing query [%s]: %v", q, err)
        }
    }
    return nil
}
