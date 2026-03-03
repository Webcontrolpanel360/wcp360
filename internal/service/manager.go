package service

import (
    "fmt"
)

// ReloadCaddy asks the system to apply new website configurations.
func ReloadCaddy() error {
    fmt.Println("[SYSTEMD] Reloading Caddy service (Simulated)...")
    return nil
}
