// A simple application for load testing how much traffic a new service can handle.
package main

import (
  "fmt"
  "net/http"
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

func http_request_worker(url_chan <-chan string) {
  for {
    url := <-url_chan
    http_get(url)
	fmt.Println("one request finished")
  }
}

func http_load_balancer(url_chan <-chan string, worker_pool_size int) {
  //create workers
  var worker_channels = make([]chan string, worker_pool_size)
  for i := range worker_channels {
    worker_channels[i] = make(chan string)
  }
  for i := 0; i < worker_pool_size; i++{
    go http_request_worker(worker_channels[i])
  }
  for {
    for _,worker_channel := range worker_channels {
	  url := <- url_chan 
	  worker_channel <- url
	}
  }
}

func main() {
  load_balancer_channel := make(chan string)
  go http_load_balancer(load_balancer_channel, 1)
  for i := 49; i < 60; i++{
    load_balancer_channel <- "http://www.example.com" 
  }
  fmt.Println("Main thread finished")
}