package ssl

import (
    "fmt"
    "os/exec"
)

// EnableHTTPS tells Caddy to provision a Let's Encrypt certificate
func EnableHTTPS(domain string) error {
    fmt.Printf("[SSL] Requesting certificate for: %s\n", domain)
    
    // Caddy handles this automatically if the domain is in the Caddyfile.
    // We force a reload to trigger the challenge.
    cmd := exec.Command("caddy", "reload", "--config", "configs/caddy/Caddyfile")
    return cmd.Run()
}
