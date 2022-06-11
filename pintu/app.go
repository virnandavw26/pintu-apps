package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", HelloServer)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, Pintu! It's Nanda :)")
}

