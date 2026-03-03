// internal/provisioner/dns/generator.go
package dns

import (
    "fmt"
    "os"
    "path/filepath"
    "time"
)

// CreateZoneFile generates a basic DNS zone for a new domain.
func CreateZoneFile(domain string, ip string) error {
    serial := time.Now().Format("2006010201")
    
    // Utilisation de fmt.Sprintf avec des verbes %s propres
    content := fmt.Sprintf("$TTL 86400\n" +
        "@   IN  SOA ns1.wcp360.com. admin.%s. (\n" +
        "        %s ; serial\n" +
        "        3600       ; refresh\n" +
        "        1800       ; retry\n" +
        "        604800     ; expire\n" +
        "        86400      ; minimum\n" +
        ")\n" +
        "@       IN  NS      ns1.wcp360.com.\n" +
        "@       IN  NS      ns2.wcp360.com.\n" +
        "@       IN  A       %s\n" +
        "www     IN  A       %s\n", domain, serial, ip, ip)

    path := filepath.Join("./configs/dns", domain + ".zone")
    return os.WriteFile(path, []byte(content), 0644)
}
