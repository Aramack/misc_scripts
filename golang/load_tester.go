// A simple application for load testing how much traffic a new service can handle.
package main

import (
  "fmt"
  "log"
  "net/http"
  "net/http/httputil"
  "net/url"
  "strings"
  "io/ioutil"
)

func http_get() {
  req := &http.Request{
          Method: "GET",
          URL: &url.URL{
            Host:   "https://example.com/",
            Scheme: "https",
          },
    	  Header: http.Header{
            "User-Agent": {"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.116 Safari/537.36"},
          },
    }
    out, err := httputil.DumpRequestOut(req, true)
    if err != nil {
      log.Fatal(err)
    }
    fmt.Println(strings.Replace(string(out), "\r", "", -1))
}

func http_get_2() {
  resp, _ := http.Get("http://example.com/")
  defer resp.Body.Close()
  body, _ := ioutil.ReadAll(resp.Body)
  fmt.Println(body)
}

func main() {
  http_get_2();
}