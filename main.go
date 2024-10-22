package main

import (
    "fmt"
    "log"
    "net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func greet(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")
    if name == "" {
        name = "Stranger"
    }
    fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {
    http.HandleFunc("/", helloWorld)
    http.HandleFunc("/greet", greet)

    log.Println("Server started on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
