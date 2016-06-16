package main

import (
  "fmt"
  "os/exec"
  "strings"
  "sync"
)

var commands = []string{"npm start"}

func exe_cmd(cmd string, wg *sync.WaitGroup) {
  // splitting head => g++ parts => rest of the command
  parts := strings.Fields(cmd)
  head := parts[0]
  parts = parts[1:len(parts)]

  out, err := exec.Command(head, parts...).Output()
  if err != nil {
    fmt.Printf("%s", err)
  }
  fmt.Printf("%s", out)
  if wg != nil {
    wg.Done()
  }
}

func main() {
  wg := new(sync.WaitGroup)
  for _, str := range commands {
    wg.Add(1)
    go exe_cmd(str, wg)
  }
  wg.Wait()
}
