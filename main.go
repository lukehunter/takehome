package main

import (
    "log"
    "net/http"
)

func main() {
    router := NewRouter()
    InitDB()

    router.Handle("/", http.FileServer(http.Dir("./static/")))
    http.Handle("/", router)

    log.Fatal(http.ListenAndServe(":8080", router))
}