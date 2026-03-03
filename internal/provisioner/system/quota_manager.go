package system

import (
    "os/exec"
    "fmt"
)

// SetUserQuota définit une limite stricte (hard limit) pour un utilisateur
// limitMB est la taille en Mégaoctets
func SetUserQuota(username string, limitMB int) error {
    // Convertir MB en blocs (1 bloc = 1KB généralement)
    blocks := limitMB * 1024
    
    // Commande setquota: setquota -u user <soft_block> <hard_block> <soft_inode> <hard_inode> /
    // On définit la limite "Hard" (blocage immédiat)
    cmd := exec.Command("setquota", "-u", username, "0", fmt.Sprintf("%d", blocks), "0", "0", "/")
    
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("impossible d'appliquer le quota: %v", err)
    }
    return nil
}
