package mail

import (
    "os"
    "fmt"
)

// CreateEmailAccount prépare le dossier mail pour Dovecot
func CreateEmailAccount(domain, emailUser string) error {
    mailPath := fmt.Sprintf("/var/vmail/%s/%s", domain, emailUser)
    
    // Créer le dossier Maildir
    err := os.MkdirAll(mailPath, 0700)
    if err != nil {
        return err
    }
    
    // En production, on ajouterait ici l'entrée dans la base de données Postfix
    fmt.Printf("[MAIL] Boîte %s@%s créée avec succès.\n", emailUser, domain)
    return nil
}
