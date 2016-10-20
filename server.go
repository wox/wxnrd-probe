package main

import (
	"io"
	"os"
	"fmt"
	"net"
	"net/http"
)

func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/hostname", Hostname)
	http.HandleFunc("/ip", Ip)
	http.ListenAndServe(":8000", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, fmt.Sprintf("Hello %s, I'm %s @ %s.\n", r.RemoteAddr, hostname(), ip()))
}

func Hostname(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, hostname())
	io.WriteString(w, "\n")
}

func Ip(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, ip())
	io.WriteString(w, "\n")
}

func hostname() string {
	hostname, _ := os.Hostname()
	return hostname	
}

func ip() string {
    addrs, err := net.InterfaceAddrs()
    if err != nil {
        return ""
    }
    for _, address := range addrs {
        if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
            if ipnet.IP.To4() != nil {
                return ipnet.IP.String()
            }
        }
    }
    return ""
}

