package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/Webcontrolpanel360/wcp360/internal/core"
	"github.com/Webcontrolpanel360/wcp360/internal/db"
	"github.com/Webcontrolpanel360/wcp360/internal/monitor"
	"github.com/Webcontrolpanel360/wcp360/internal/security/waf"
	"github.com/Webcontrolpanel360/wcp360/internal/security/ssl"
	"github.com/Webcontrolpanel360/wcp360/internal/terminal"
	"github.com/Webcontrolpanel360/wcp360/internal/diag"
	"github.com/Webcontrolpanel360/wcp360/internal/backup"
)

func main() {
	// 1. Initialize DB
	db.InitDB()
	
	// 2. Initialize Engine
	_ = core.NewEngine("2.2.0")
	
	// 3. Setup Echo
	e := echo.New()

	// 4. Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(waf.Middleware())

	// 5. Basic Auth for Admin Panel
	e.Use(middleware.BasicAuth(func(u, p string, c echo.Context) (bool, error) {
		return u == "admin" && p == "wcp360-secure", nil
	}))

	// 6. Static Routes
	e.File("/", "ui/templates/index.html")

	// 7. API Routes
	api := e.Group("/api")
	{
		api.GET("/system", func(c echo.Context) error { 
			return c.JSON(http.StatusOK, monitor.GetStats()) 
		})
		
		api.GET("/health", func(c echo.Context) error { 
			return c.JSON(http.StatusOK, diag.CheckSystem()) 
		})

		api.POST("/ssl/enable", func(c echo.Context) error {
			req := struct { Domain string `json:"domain"` }{}
			if err := c.Bind(&req); err != nil { return err }
			ssl.EnableHTTPS(req.Domain)
			return c.JSON(http.StatusOK, map[string]string{"message": "SSL request sent"})
		})

		api.POST("/db/create", func(c echo.Context) error {
			req := struct { DBName, User, Pass string }{}
			if err := c.Bind(&req); err != nil { return err }
			return c.JSON(http.StatusCreated, map[string]string{"status": "Database request received"})
		})
		api.POST("/backup/create", func(c echo.Context) error {
            req := struct { User, DB string }{ }
            c.Bind(&req)
            // Lancer le backup en arrière-plan (Go routine)
            go backup.CreateFullBackup(req.User, req.DB)
            return c.JSON(http.StatusAccepted, map[string]string{"message": "Backup started"})
        })
		// Installer WordPress
        api.POST("/apps/install/wordpress", func(c echo.Context) error {
            req := struct {
                User, Domain, DBName, DBUser, DBPass string
            }{}
            c.Bind(&req)
            err := apps.InstallWordPress(req.User, req.Domain, req.DBName, req.DBUser, req.DBPass)
            if err != nil { return c.JSON(500, err.Error()) }
            return c.JSON(200, "WordPress installé !")
        })

        // Lister les fichiers pour l'explorateur
        api.GET("/files/list", func(c echo.Context) error {
            path := c.QueryParam("path") // ex: /home/user/public_html
            files, _ := filesystem.ListFiles(path)
            return c.JSON(200, files)
        })
	}

	// 8. WebSocket for Terminal
	e.GET("/ws/terminal", func(c echo.Context) error { 
		terminal.HandleTerminal(c.Response(), c.Request())
		return nil 
	})

	// 9. Start Server
	e.Logger.Fatal(e.Start(":8080"))
}
