// internal/quota/manager.go
// This module calculates disk usage for each tenant.
package quota

import (
    "os"
    "path/filepath"
)

// GetDirSize calculates the total size of a directory in bytes.
func GetDirSize(username string) (int64, error) {
    path := filepath.Join("./data/www", username)
    var size int64
    err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.IsDir() {
            size += info.Size()
        }
        return err
    })
    return size, err
}
