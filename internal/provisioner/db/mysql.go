package db

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

// CreateDatabase crée une base, un utilisateur et donne les accès.
func CreateDatabase(dbName, dbUser, dbPass string) error {
    // Connexion au MySQL local (root)
    // En prod, ces accès seront dans un fichier de config
    dsn := "root:password@tcp(127.0.0.1:3306)/"
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return err
    }
    defer db.Close()

    // Liste des commandes SQL à exécuter
    queries := []string{
        fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName),
        fmt.Sprintf("CREATE USER IF NOT EXISTS '%s'@'localhost' IDENTIFIED BY '%s'", dbUser, dbPass),
        fmt.Sprintf("GRANT ALL PRIVILEGES ON %s.* TO '%s'@'localhost'", dbName, dbUser),
        "FLUSH PRIVILEGES",
    }

    for _, q := range queries {
        if _, err := db.Exec(q); err != nil {
            return fmt.Errorf("Erreur SQL [%s]: %v", q, err)
        }
    }
    return nil
}
