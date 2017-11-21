package main

import (
    "os"
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, this is go supported by custom buildpack.")
}

func main() {
    port := os.Getenv("PORT")
    if port == "" {
      port = "3000"
    }
    http.HandleFunc("/", handler)
    fmt.Printf("Server running on http://localhost:%s ...\n", port)
    http.ListenAndServe(fmt.Sprintf(":" + port), nil)
}
