// A simple application for load testing how much traffic a new service can handle.
package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "time"
)

func http_get(url string) {
  request , err := http.NewRequest(
    "GET",
	url,
	nil,
  )
  if (err != nil) {
    return
  }
  client := &http.Client{
    Timeout: time.Second * 30,
  }
  client.Do(request)
}

func http_get_2() {
  resp, _ := http.Get("http://example.com/")
  defer resp.Body.Close()
  body, _ := ioutil.ReadAll(resp.Body)
  fmt.Println(body)
}

func main() {
  //http_get_2();
  http_get("http://www.example.com")
}