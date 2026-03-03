package provisioner

import (
    "fmt"
    "os"
    "path/filepath"
    "github.com/Webcontrolpanel360/wcp360/internal/models"
    "github.com/Webcontrolpanel360/wcp360/internal/service"
)

func SetupWebsite(t models.Tenant) error {
    basePath := fmt.Sprintf("./data/www/%s", t.Username)
    configPath := fmt.Sprintf("./configs/caddy/%s.caddy", t.Username)

    os.MkdirAll(basePath, 0755)
    indexContent := fmt.Sprintf("<h1>Welcome to %s</h1>", t.Domain)
    os.WriteFile(filepath.Join(basePath, "index.html"), []byte(indexContent), 0644)

    // Notice: We removed HTTP/3 and manual SSL because Caddy does it automatically
    // just by putting the domain name at the top.
    caddyConfig := fmt.Sprintf("%s {\n    root * %s\n    file_server\n    encode zstd gzip\n}\n", t.Domain, basePath)
    
    os.WriteFile(configPath, []byte(caddyConfig), 0644)

    // Trigger a reload so the new site works immediately
    service.ReloadCaddy()
    
    return nil
}
