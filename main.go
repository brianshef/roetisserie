package main

import (
    "fmt"
    "log"
    "net/http"
    "os/exec"
    "runtime"
    "time"

    "github.com/brianshef/roetisserie/server"
)

var host = "http://localhost"
var port = "3000"
var url = host + ":" + port
var healthcheck_endpoint = "/healthcheck"

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
    log.Print("OS: ", runtime.GOOS)

    //  Client routine
    go func () {
        for {
            time.Sleep(time.Second)
            log.Println("Client checking for server connection ... ")
            resp, err := http.Get(url + healthcheck_endpoint)
            if err != nil {
                log.Println("Failed:", err)
                continue
            }

            resp.Body.Close()
            if resp.StatusCode != http.StatusOK {
                continue
            }

            // Server returned 200 OK
            break
        }

        log.Print("Client connection successful --> ", url)
        open_browser(url)
    }()

    //  Server blocks without error; it must be started last, AFTER client
    server.Start()
}
