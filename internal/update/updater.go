package update

import (
    "os/exec"
    "fmt"
)

// CheckAndPull connects to GitHub and pulls the latest main branch
func CheckAndPull() (string, error) {
    fmt.Println("[UPDATER] Checking for updates...")
    
    // In a real scenario, you'd check tags via GitHub API
    // For now, we perform a git pull
    cmd := exec.Command("git", "pull", "origin", "main")
    out, err := cmd.CombinedOutput()
    if err != nil {
        return "", err
    }
    return string(out), nil
}
