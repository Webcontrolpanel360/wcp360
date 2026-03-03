// internal/logs/viewer.go
// This module reads the last lines of a log file to help debugging.
package logs

import (
    "os/exec"
    "fmt"
)

// GetLastLogs reads the last 50 lines of a specific log file using the 'tail' command.
func GetLastLogs(logType string) (string, error) {
    // In production, you would point to /var/log/caddy/access.log
    // For now, we simulate by reading our own panel output or a dummy file.
    fmt.Printf("[LOGS] Fetching last entries for: %s\n", logType)
    
    cmd := exec.Command("tail", "-n", "50", "wcp360.db") // Reading DB as dummy text for test
    out, err := cmd.CombinedOutput()
    if err != nil {
        return "No logs available yet.", nil
    }
    return string(out), nil
}
