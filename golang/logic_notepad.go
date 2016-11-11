package main

import "fmt"

func main() {
  a := make(map[string]bool)
  a["test"] = false
  if _, ok := a["test"]; !ok {
    fmt.Println("true")
  } else {
    fmt.Println("false")
  }
}
