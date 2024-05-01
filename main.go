package main

import (
	"fmt"
	"net"
	"net/http"
)

func headers(w http.ResponseWriter, req *http.Request) {
	
	host, _, _ := net.SplitHostPort(req.RemoteAddr)
	req.Header.Set("X-Forwarded-For", host)

	for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func main() {
    http.HandleFunc("/solicitar/comprovante/d2c4bb01-5327-40fd-aeba-e11416c28f63.html", headers)

    http.ListenAndServe(":8090", nil)
}