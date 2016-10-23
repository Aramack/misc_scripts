// A simple application for load testing how much traffic a new service can handle.
package main

import (
  "fmt"
  "net/http"
  "time"
  "os"
  "bufio"
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
    fmt.Println("one request finished: " + url)
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

func read_url_source(raw_url_chan chan<- string, file_path string) {
  file_handle, _ := os.Open(file_path)
  reader := bufio.NewScanner(file_handle)
  reader.Split(bufio.ScanLines)
  
  for reader.Scan() {
    raw_url_chan <- reader.Text()
  }
  file_handle.Close()
  close(raw_url_chan)
}

func main() {
  load_balancer_channel := make(chan string)
  go http_load_balancer(load_balancer_channel, 4)

  raw_url_chan := make(chan string)
  go read_url_source(raw_url_chan, "/tmp/test_url")
  
  for {
    raw_url, more := <- raw_url_chan
    if more {
      load_balancer_channel <- raw_url
    } else {
      break
    }
  }
  fmt.Println("Main thread finished")
}