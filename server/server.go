package server

import (
    "net/http"
    "log"
)

var port string = "3000"

func serve_main(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "client/main.html")
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("OK"))
}

func Start() {
    http.HandleFunc("/", serve_main)
    http.HandleFunc("/healthcheck", healthcheck)

    log.Print("Server listening on port ", port)
    err := http.ListenAndServe(":" + port, nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
