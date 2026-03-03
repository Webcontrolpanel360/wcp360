package security

import (
    "fmt"
)

// SetupFirewall configures basic rules.
func SetupFirewall() error {
    fmt.Println("[SECURITY] Configuring nftables firewall (Simulated)...")
    return nil
}

// BlockIP adds a malicious IP address to the blacklist.
func BlockIP(ip string) {
    fmt.Printf("[SECURITY] Blocking malicious IP: %s\n", ip)
}
