# 🐹 Go REST API — Built From Scratch

A simple REST API built with **pure Go** — no frameworks, no extra libraries. Just Go's built-in `net/http` package.


---

## 🎬 Watch the Video

If you landed here from the video — hey! 👋 The full code is right here.

---

## 🚀 What This API Does

| Route | Method | What it returns |
|-------|--------|-----------------|
| `/health` | GET | Confirms the server is running |
| `/users` | GET | Returns a list of users as JSON |

---

## 🛠️ Requirements

- [Go](https://go.dev/dl/) version 1.21 or higher
- That's it — no npm, no pip, nothing else

Check your Go version by running:
```bash
go version
```

---

## ⚡ Quick Start

**1. Clone the repo**
```bash
git clone https://github.com/ghosthouse7/go-rest-api
cd go-rest-api
```

**2. Run the server**
```bash
go run .
```

You should see:
```
Server starting on port 8080...
```

**3. Test it**

Open your browser or Postman and hit these URLs:

```
GET http://localhost:8080/health
GET http://localhost:8080/users
```

---

## 📦 Expected Responses

**GET /health**
```
API is running
```

**GET /users**
```json
[
  {"id": 1, "name": "Alice", "email": "alice@example.com"},
  {"id": 2, "name": "Bob",   "email": "bob@example.com"},
  {"id": 3, "name": "Charlie", "email": "charlie@example.com"}
]
```

---

## 🗂️ Project Structure

```
go-rest-api/
│
├── main.go        ← Everything lives here (for now!)
└── go.mod         ← Go module file
```

> Yes, it's just one file. That's the point — keep it simple, understand every line.

---

## 🧠 What You'll Learn From This Code

- How to create an HTTP server in Go using `net/http`
- How to define routes with `http.HandleFunc`
- How to use **structs** to model your data
- How to return **JSON responses** using `encoding/json`
- How to check HTTP methods (GET, POST, etc.)
- Why `log.Fatal` matters for error handling

---

## 📁 The Full Code

```go
package main

import (
    "encoding/json"
    "log"
    "net/http"
)

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

var users = []User{
    {ID: 1, Name: "Alice",   Email: "alice@example.com"},
    {ID: 2, Name: "Bob",     Email: "bob@example.com"},
    {ID: 3, Name: "Charlie", Email: "charlie@example.com"},
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("API is running"))
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

func main() {
    http.HandleFunc("/health", healthHandler)
    http.HandleFunc("/users", getUsersHandler)

    log.Println("Server starting on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

---

## 🔮 What's Coming Next

This is Part 1. Here's what we're building next on the channel:

- [ ] **Part 2** — POST /users endpoint to create new users
- [ ] **Part 3** — Connect to a real PostgreSQL database
- [ ] **Part 4** — Add JWT authentication
- [ ] **Part 5** — Dockerize the whole thing

STAY TUNED FOR MORE!!!!

---

## 🤝 Contributing

Found a bug or want to suggest something? Open an issue or drop a comment on the video. This repo is meant to be a learning resource so all feedback is welcome!

---
<p align="center">Made with ☕ and Go 
