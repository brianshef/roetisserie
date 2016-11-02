package main

import (
    "fmt"
    "log"
    "os/exec"
    "runtime"
    "github.com/brianshef/roetisserie/server"
)

var url = "http://localhost:3000"

func open_browser(url string) {
    log.Print("Opening browser ... ")
    var err error
    switch runtime.GOOS {
    case "linux":
        err = exec.Command("xdg-open", url).Start()
    case "windows", "darwin":
        err = exec.Command("open", url).Start()
    default:
        err = fmt.Errorf("unsupported platform")
    }

    if err != nil {
        log.Fatal(err)
    }
}

func main() {
    log.Print(runtime.GOOS)
    open_browser(url)
    server.Start()
}
