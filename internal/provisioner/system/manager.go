package system

import (
	"path/filepath"
	"strings"
)

// SafePath ensures the file path is within the user's home directory
func SafePath(username, filename string) string {
	root := "/tmp/wcp360_home/" + username + "/admin"
	path := filepath.Join(root, filename)
	if !strings.HasPrefix(path, root) {
		return root // Default to root if directory traversal is attempted
	}
	return path
}
