// internal/db/database.go
// This file handles the SQLite connection to store tenant data.
package db

import (
    "github.com/Webcontrolpanel360/wcp360/internal/models"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
    var err error
    // Connect to SQLite file defined in config
    DB, err = gorm.Open(sqlite.Open("wcp360.db"), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database!")
    }
    // Automatically create the 'tenants' table based on our model
    DB.AutoMigrate(&models.Tenant{})
}
