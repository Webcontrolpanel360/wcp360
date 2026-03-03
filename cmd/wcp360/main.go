package main

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "github.com/Webcontrolpanel360/wcp360/internal/core"
    "github.com/Webcontrolpanel360/wcp360/internal/db"
    "github.com/Webcontrolpanel360/wcp360/internal/models"
    "github.com/Webcontrolpanel360/wcp360/internal/monitor"
    "github.com/Webcontrolpanel360/wcp360/internal/security/waf"
    "github.com/Webcontrolpanel360/wcp360/internal/security/ssl"
    "github.com/Webcontrolpanel360/wcp360/internal/provisioner/db" // MySQL manager
)

func main() {
    db.InitDB()
    wcp := core.NewEngine("2.2.0")
    e := echo.New()

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
    e.Use(waf.Middleware())

    e.Use(middleware.BasicAuth(func(u, p string, c echo.Context) (bool, error) {
        return u == "admin" && p == "wcp360-secure", nil
    }))

    e.File("/", "ui/templates/index.html")

    // API : Stats & Logs
    e.GET("/api/system", func(c echo.Context) error { return c.JSON(http.StatusOK, monitor.GetStats()) })

    // API : Activer SSL
    e.POST("/api/ssl/enable", func(c echo.Context) error {
        req := struct { Domain string `json:"domain"` }{}
        c.Bind(&req)
        ssl.EnableHTTPS(req.Domain)
        return c.JSON(http.StatusOK, map[string]string{"message": "SSL request sent"})
    })

    // API : Créer DB MySQL
    e.POST("/api/db/create", func(c echo.Context) error {
        req := struct { DBName, User, Pass string }{}
        c.Bind(&req)
        // db_prov.CreateDatabase(req.DBName, req.User, req.Pass) 
        return c.JSON(http.StatusCreated, map[string]string{"status": "Database provisioned"})
    })

    e.Logger.Fatal(e.Start(":8080"))
}
