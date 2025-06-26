package terminal

import (
	"log"
	"net/http"
	"os/exec"

	"github.com/creack/pty"
	"github.com/gorilla/websocket"
)

var runCommand []string

func SetCommand(cmd []string) {
	runCommand = cmd
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func Handler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade error: ", err)
		return
	}
	defer ws.Close()

	cmd := exec.Command(runCommand[0], runCommand[1:]...)

	ptmx, err := pty.Start(cmd)
	if err != nil {
		log.Println("PTY start error: ", err)
		return
	}
	defer ptmx.Close()

	// PTY → WebSocket
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := ptmx.Read(buf)
			if err != nil {
				break
			}
			ws.WriteMessage(websocket.TextMessage, buf[:n])
		}
	}()

	// WebSocket → PTY
	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			break
		}
		ptmx.Write(msg)
	}
}
