package main

import (
  "log"
  // "net/http"
  "github.com/armon/go-socks5"
)

func main() {
  // Create a SOCKS5 server
  conf := &socks5.Config{}
  server, err := socks5.New(conf)
  if err != nil {
    panic(err)
  }

  // Create SOCKS5 proxy on localhost port 8000
  log.Printf("sock5 proxy listen on: %s", "8080")

  if err := server.ListenAndServe("tcp", "0.0.0.0:8080"); err != nil {
    panic(err)
  }
}