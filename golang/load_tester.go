// A simple application for load testing how much traffic a new service can handle.
package main

import (
  "fmt"
  "log"
  "net/http"
  "net/http/httputil"
  "net/url"
  "strings"
)

func http_get() {
  req := &http.Request{
          Method: "GET",
          Host:   "https://example.com/",
          URL: &url.URL{
            Host:   "ignored",
            Scheme: "https",
            Opaque: "/%2f/",
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

func main() {
  http_get();
}