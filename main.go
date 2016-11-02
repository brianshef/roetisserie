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

var url = "http://localhost:3000"
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

    go func () {
        check_url := url + healthcheck_endpoint
        for {
            time.Sleep(time.Second)
            log.Println("Checking for server connection ... ")
            resp, err := http.Get(check_url)
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

        log.Print("SERVER UP AND RUNNING @ ", url)
        open_browser(url)
    }()

    server.Start()
}
