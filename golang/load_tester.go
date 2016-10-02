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
  resp, err := client.Do(request)
  resp.Body.Close()
}

func http_get_2() {
  resp, _ := http.Get("http://example.com/")
  defer resp.Body.Close()
  body, _ := ioutil.ReadAll(resp.Body)
  fmt.Println(body)
}

func http_request_worker(url_chan <-chan string) {
  
  
  url := <-url_chan
  http_get(url)
  

}

func main() {
  //http_get_2();
  url_chan := make(chan string)
  go http_request_worker(url_chan)
  url_chan <- "http://www.example.com" 
  
}