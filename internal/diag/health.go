package diag

import (
	"os/exec"
	"runtime"
)

type HealthStatus struct {
	OS      string `json:"os"`
	MySQL   bool   `json:"mysql_running"`
	Caddy   bool   `json:"caddy_running"`
	Storage string `json:"storage_status"`
}

func CheckSystem() HealthStatus {
	// Vérifie si MySQL ou MariaDB tourne
	mysqlCheck := exec.Command("pgrep", "mysql").Run() == nil || exec.Command("pgrep", "mariadbd").Run() == nil
	
	// Vérifie si Caddy tourne
	caddyCheck := exec.Command("pgrep", "caddy").Run() == nil

	return HealthStatus{
		OS:      runtime.GOOS,
		MySQL:   mysqlCheck,
		Caddy:   caddyCheck,
		Storage: "Healthy",
	}
}
