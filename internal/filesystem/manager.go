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

// ListFiles retourne la liste des fichiers et dossiers avec leurs détails
type FileInfo struct {
    Name string `json:"name"`
    Size int64  `json:"size"`
    IsDir bool  `json:"is_dir"`
    Time string `json:"time"`
}

func ListFiles(path string) ([]FileInfo, error) {
    entries, err := os.ReadDir(path)
    if err != nil { return nil, err }

    var files []FileInfo
    for _, entry := range entries {
        info, _ := entry.Info()
        files = append(files, FileInfo{
            Name:  entry.Name(),
            Size:  info.Size(),
            IsDir: entry.IsDir(),
            Time:  info.ModTime().Format("2006-01-02 15:04"),
        })
    }
    return files, nil
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

// ReadFileContent lit le contenu d'un fichier texte
func ReadFileContent(username, filename string) (string, error) {
    path := filepath.Join("/var/www/wcp360/data/www", username, filename)
    content, err := os.ReadFile(path)
    return string(content), err
}

// SaveFileContent écrit le contenu dans un fichier
func SaveFileContent(username, filename, content string) error {
    path := filepath.Join("/var/www/wcp360/data/www", username, filename)
    return os.WriteFile(path, []byte(content), 0644)
}
