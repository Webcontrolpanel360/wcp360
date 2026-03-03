package main

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "github.com/Webcontrolpanel360/wcp360/internal/core"
    "github.com/Webcontrolpanel360/wcp360/internal/db"
    "github.com/Webcontrolpanel360/wcp360/internal/models"
    "github.com/Webcontrolpanel360/wcp360/internal/monitor"
    "github.com/Webcontrolpanel360/wcp360/internal/filesystem"
    "github.com/Webcontrolpanel360/wcp360/internal/security/waf"
    "github.com/Webcontrolpanel360/wcp360/internal/backup"
    "github.com/Webcontrolpanel360/wcp360/internal/logs"
    "github.com/Webcontrolpanel360/wcp360/internal/provisioner/php"
    "github.com/Webcontrolpanel360/wcp360/internal/quota"
    "github.com/Webcontrolpanel360/wcp360/internal/provisioner/dns"
)

func main() {
    db.InitDB()
    wcp := core.NewEngine("2.0.0")
    e := echo.New()

    // --- MIDDLEWARES ---
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
    e.Use(waf.Middleware())

    // --- AUTHENTICATION ---
    e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
        if username == "admin" && password == "wcp360-secure" {
            return true, nil
        }
        return false, nil
    }))

    // --- ROUTES ---
    e.File("/", "ui/templates/index.html")

    // API DNS (Utilise maintenant le module dns)
    e.POST("/api/dns/generate", func(c echo.Context) error {
        req := struct { Domain string `json:"domain"`; IP string `json:"ip"` }{}
        if err := c.Bind(&req); err != nil {
            return err
        }
        dns.CreateZoneFile(req.Domain, req.IP)
        return c.JSON(http.StatusOK, map[string]string{"message": "DNS Zone created"})
    })

    e.GET("/api/system", func(c echo.Context) error { return c.JSON(http.StatusOK, monitor.GetStats()) })
    
    e.GET("/api/tenants", func(c echo.Context) error {
        var tenants []models.Tenant
        db.DB.Find(&tenants)
        return c.JSON(http.StatusOK, tenants)
    })

    e.POST("/api/tenants", func(c echo.Context) error {
        t := new(models.Tenant)
        c.Bind(t); db.DB.Create(&t); wcp.ProvisionTenant(*t)
        return c.JSON(http.StatusCreated, t)
    })

    e.GET("/api/files/:username", func(c echo.Context) error {
        user := c.Param("username"); files, _ := filesystem.ListFiles(user)
        return c.JSON(http.StatusOK, files)
    })

    e.POST("/api/backup/:username", func(c echo.Context) error {
        user := c.Param("username"); file, _ := backup.CreateBackup(user)
        return c.JSON(http.StatusOK, map[string]string{"file": file})
    })

    e.POST("/api/php/set", func(c echo.Context) error {
        req := struct { User string; Version string }{}; c.Bind(&req)
        php.SetVersion(req.User, req.Version)
        return c.JSON(http.StatusOK, map[string]string{"message": "PHP updated"})
    })

    e.GET("/api/quota/:username", func(c echo.Context) error {
        user := c.Param("username"); size, _ := quota.GetDirSize(user)
        return c.JSON(http.StatusOK, map[string]interface{}{"size_mb": float64(size)/1024/1024})
    })

    e.GET("/api/logs", func(c echo.Context) error {
        data, _ := logs.GetLastLogs("system")
        return c.JSON(http.StatusOK, map[string]string{"logs": data})
    })

    e.Logger.Fatal(e.Start(":8080"))
}
