package main

import (
  "fmt"
  "os/exec"
  "runtime"
  "./server"
)

var url = "http://localhost:3000"

func open_browser(url string) {
  fmt.Println("Opening browser ... ")
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
    fmt.Println(err)
  }
}

func main() {
    fmt.Println(runtime.GOOS)

    server.Start()

    open_browser(url)
}
