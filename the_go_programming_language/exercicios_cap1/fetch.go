package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
  "os"
  "strings"
)

func main() {
  for _, url := range os.Args[1:] {
    if !(strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")){ 
      url = "http://" + url
    }
    resp, err := http.Get(url)
    if err != nil {
      fmt.Fprintf(os.Stderr, "fetch: %v : %s\n", err,resp.Status)
      os.Exit(1)
    }
    b, err := ioutil.ReadAll(resp.Body)
    resp.Body.Close()
    if err != nil {
      fmt.Fprintf(os.Stderr, "fetch: reading %s: %v : %s\n", url, err,resp.Status)
      os.Exit(1)
    }
    fmt.Printf("%s \nStatus code - %s\n", b,resp.Status)
  }
}
