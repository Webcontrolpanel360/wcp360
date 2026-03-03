package cron

import (
    "os/exec"
    "fmt"
    "os"
)

// AddCronJob ajoute une ligne au crontab de l'utilisateur root
func AddCronJob(schedule, command string) error {
    // On récupère le crontab actuel
    current, _ := exec.Command("crontab", "-l").Output()
    
    newCron := string(current) + fmt.Sprintf("%s %s\n", schedule, command)
    
    // On écrit le nouveau crontab dans un fichier temporaire
    tmpFile := "/tmp/new_cron"
    os.WriteFile(tmpFile, []byte(newCron), 0644)
    
    // On applique le nouveau crontab
    cmd := exec.Command("crontab", tmpFile)
    return cmd.Run()
}
