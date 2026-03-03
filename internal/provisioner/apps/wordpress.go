package apps

import (
    "fmt"
    "os/exec"
    "os"
)

// InstallWordPress automatise le déploiement de WP
func InstallWordPress(username, domain, dbName, dbUser, dbPass string) error {
    path := fmt.Sprintf("/home/%s/public_html", username)

    // 1. Téléchargement de la dernière version
    exec.Command("wget", "https://wordpress.org/latest.tar.gz", "-P", "/tmp").Run()

    // 2. Extraction
    exec.Command("tar", "-xzvf", "/tmp/latest.tar.gz", "-C", "/tmp").Run()
    exec.Command("cp", "-R", "/tmp/wordpress/.", path).Run()

    // 3. Configuration des permissions
    exec.Command("chown", "-R", username+":"+username, path).Run()

    // 4. Nettoyage
    os.Remove("/tmp/latest.tar.gz")
    os.RemoveAll("/tmp/wordpress")

    fmt.Printf("[APP] WordPress installé pour %s sur %s\n", username, domain)
    return nil
}
