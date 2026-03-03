// internal/provisioner/php/selector.go
// This module manages different PHP versions for tenants.
package php

import (
    "fmt"
)

// SetVersion updates the PHP handler for a specific tenant.
func SetVersion(username string, version string) error {
    // Valid versions: 8.1, 8.2, 8.3
    fmt.Printf("[PHP] Switching user %s to PHP %s\n", username, version)
    
    // In production, this would update the Caddyfile or change the 
    // upstream socket to /var/run/php/php${version}-fpm.sock
    return nil
}

func GetAvailableVersions() []string {
    return []string{"8.1", "8.2", "8.3"}
}
