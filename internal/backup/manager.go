// internal/backup/manager.go
// This module handles the creation of site backups.
package backup

import (
    "fmt"
    "os/exec"
    "time"
)

// CreateBackup creates a compressed archive of a tenant's web directory.
func CreateBackup(username string) (string, error) {
    timestamp := time.Now().Format("20060102-150405")
    backupFile := fmt.Sprintf("./backups/%s-%s.tar.gz", username, timestamp)
    sourceDir := fmt.Sprintf("./data/www/%s", username)

    // Execute the Linux 'tar' command to compress the folder
    cmd := exec.Command("tar", "-czf", backupFile, sourceDir)
    err := cmd.Run()
    
    if err != nil {
        return "", fmt.Errorf("failed to create backup: %w", err)
    }

    fmt.Printf("[BACKUP] Created: %s\n", backupFile)
    return backupFile, nil
}
