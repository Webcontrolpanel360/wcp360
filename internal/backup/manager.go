package backup

import (
    "fmt"
    "os/exec"
    "time"
)

// CreateFullBackup compresse le dossier d'un utilisateur et exporte sa base SQL
func CreateFullBackup(username, dbName string) (string, error) {
    timestamp := time.Now().Format("2006-01-02_15-04")
    backupPath := fmt.Sprintf("/home/%s/backups/backup_%s.tar.gz", username, timestamp)
    
    // 1. Export MySQL (Dump)
    // mysqldump -u root [DB_NAME] > /tmp/[DB_NAME].sql
    sqlPath := fmt.Sprintf("/tmp/%s.sql", dbName)
    exec.Command("mysqldump", dbName, "-r", sqlPath).Run()

    // 2. Compression du dossier public_html + le dump SQL
    // tar -czf [DEST] [SOURCE]
    cmd := exec.Command("tar", "-czf", backupPath, "-C", "/home/"+username, "public_html", "-C", "/tmp", dbName+".sql")
    
    err := cmd.Run()
    return backupPath, err
}
