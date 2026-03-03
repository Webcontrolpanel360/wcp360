package terminal

import (
	"net/http"
	"os/exec"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

// wsWrapper permet de lier WebSocket à io.Reader et io.Writer
type wsWrapper struct {
	conn *websocket.Conn
}

// Read implémente io.Reader : lit les messages du navigateur et les envoie au shell
func (w *wsWrapper) Read(p []byte) (n int, err error) {
	_, message, err := w.conn.ReadMessage()
	if err != nil {
		return 0, err
	}
	return copy(p, message), nil
}

// Write implémente io.Writer : prend la sortie du shell et l'envoie au navigateur
func (w *wsWrapper) Write(p []byte) (n int, err error) {
	err = w.conn.WriteMessage(websocket.TextMessage, p)
	if err != nil {
		return 0, err
	}
	return len(p), nil
}

func HandleTerminal(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	// Lancer un shell Bash
	cmd := exec.Command("bash")
	
	// Créer un wrapper pour notre connexion
	ws := &wsWrapper{conn: conn}

	// Lier les entrées/sorties du shell au WebSocket
	cmd.Stdout = ws
	cmd.Stderr = ws
	cmd.Stdin = ws

	err = cmd.Run()
}
