package filesystem

import (
    "os"
    "path/filepath"
)

type FileInfo struct {
    Name  string `json:"name"`
    IsDir bool   `json:"is_dir"`
    Size  int64  `json:"size"`
}

// ListFiles lists content of a tenant's web directory
func ListFiles(username string) ([]FileInfo, error) {
    root := filepath.Join("./data/www", username)
    files, err := os.ReadDir(root)
    if err != nil {
        return nil, err
    }

    var result []FileInfo
    for _, f := range files {
        info, _ := f.Info()
        result = append(result, FileInfo{
            Name:  f.Name(),
            IsDir: f.IsDir(),
            Size:  info.Size(),
        })
    }
    return result, nil
}
