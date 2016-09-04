package main

import (
  "fmt"
  "time"
)

func main() {
  for i := 0; i < 5; i++ {
    go do_something(i)
  }
  time.Sleep(1 * time.Second)
}

func do_something(i int) {
  time.Sleep(100 * time.Millisecond)
  fmt.Println(i)
}
