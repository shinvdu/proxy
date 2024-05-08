package main

import (
    "github.com/elazarl/goproxy"
    "log"
    "net/http"
)

func main() {
    proxy := goproxy.NewProxyHttpServer()
    proxy.Verbose = true
    log.Printf("http proxy listen on: %s", "8080")
    log.Fatal(http.ListenAndServe(":8181", proxy))
}
