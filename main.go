package main

import (
  "fmt"
  "os/exec"
  "runtime"
  "strings"
  "sync"
)

var url = "http://localhost:3000"
var commands = []string{"npm start"}

func exe_cmd(cmd string, wg *sync.WaitGroup) {
  // splitting head => g++ parts => rest of the command
  parts := strings.Fields(cmd)
  head := parts[0]
  parts = parts[1:len(parts)]

  fmt.Println("Executing command", cmd, "...")

  out, err := exec.Command(head, parts...).Output()
  if err != nil {
    fmt.Printf("%s", err)
  }
  fmt.Printf("%s", out)
  if wg != nil {
    wg.Done()
  }
}

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
  wg := new(sync.WaitGroup)
  for _, str := range commands {
    wg.Add(1)
    go exe_cmd(str, wg)
  }
  wg.Wait()
  // open_browser(url)
}
