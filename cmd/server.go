package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/javy99/termview/internal/terminal"
)

func Run() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: termview <command> [args...]")
		os.Exit(1)
	}

	terminal.SetCommand(os.Args[1:])

	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/ws", terminal.Handler)

	port := "8080"
	log.Printf("Serving %v at http://localhost:%s\n", os.Args[1:], port)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, nil))
}
