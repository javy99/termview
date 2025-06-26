# 🖥️ termview

![Go](https://img.shields.io/badge/Go-1.24-blue)
![License](https://img.shields.io/badge/license-MIT-green)

> Share your terminal in the browser using Go, PTY, and WebSockets — works with any CLI app like Bash, Python, Go, etc.

---

## 🚀 Quick Start

### 1. Build the binary

```bash
go build -o termview
```

### 2. Run any terminal-based app or script

```bash
# Share Bash
./termview bash

# Share a Python REPL
./termview python3

# Share a Python script
./termview python3 ./scripts/hello.py

# Run a Go program
./termview go run myapp.go

# Share any binary or CLI tool
./termview ./my-cli-tool
```

✅ You can run anything that works in a terminal — Bash, Python, Node.js, REPLs, scripts, compiled apps, and more!

### 3. Open in browser

Visit: http://localhost:8080
Use your terminal app live from the browser!

## 📦 What termview Does

termview lets you share any terminal-based application over the web.
  - Run any CLI command or REPL (e.g. bash, python3, go run main.go)
  - View and interact with it inside your web browser
  - Works across devices (even mobile browsers)
  - Ideal for teaching, debugging, demos, or remote terminal sessions

## 🔧 Tech Stack

| Component       | What It Does                                                |
|----------------|-------------------------------------------------------------|
| **Go**         | Backend logic and WebSocket server                          |
| **PTY**        | Creates a fake terminal so apps like `bash` behave normally |
| **xterm.js**   | Browser-based terminal emulator                             |
| **WebSockets** | Two-way live connection between browser and Go server       |
| **FileServer** | Serves frontend assets (`index.html`, JS, CSS)              |


## 🔄 Architecture Diagram

```text
+--------------------------+
|     Your Web Browser     |
|  (xterm.js terminal UI)  |
+------------+-------------+
             |
             | WebSocket (keyboard ↔ output)
             v
+------------+-------------+
|    Go WebSocket Server   |
|     (termview backend)   |
+------------+-------------+
             |
             | PTY (fake terminal device)
             v
+--------------------------+
|   CLI App (bash, python) |
+--------------------------+
```

## 📁 Project Structure
```text
termview/
├── main.go                  # App entry point
├── cmd/
│   └── server.go            # Server setup and flags
├── internal/
│   └── terminal/
│       └── handler.go       # WebSocket + PTY handling logic
├── static/
│   └── index.html           # Frontend terminal UI with xterm.js
└── README.md
```

## 📌 Use Cases
  - Live coding interviews
  - Teaching Python/Bash/Go remotely
  - Remote debugging
  - Web-based REPL
  - Sharing CLI-only tools

## 📜 License

This project is licensed under the [MIT License](LICENSE).


## 🤝 Contributing

Contributions, issues, and feature requests are welcome!  
Feel free to open an issue or pull request.

---

## 🛠️ TODO

- [ ] Add password protection
- [ ] Support remote session recording
- [ ] Docker support
- [ ] Multiple user sessions