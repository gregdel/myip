package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type app struct {
	ip     string
	port   string
	logger *log.Logger
}

func newApp(ip, port string) *app {
	return &app{
		ip:     ip,
		port:   port,
		logger: log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds),
	}
}

func (a *app) publicIPHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	ip := r.Header.Get("X-Remote-Ip")
	if ip == "" {
		ip = lastIndexIP(r.RemoteAddr)
	}
	io.WriteString(w, ip)
	a.logger.Printf("ip=%s elapsed=%s", ip, time.Since(start))
}

func (a *app) Run() error {
	listenAddr := fmt.Sprintf("%s:%s", a.ip, a.port)
	h := &http.Server{
		Addr:           listenAddr,
		IdleTimeout:    0,
		ReadTimeout:    1 * time.Second,
		WriteTimeout:   1 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	http.HandleFunc("/", a.publicIPHandler)
	fmt.Printf("HTTP server started on %s\n", listenAddr)
	return h.ListenAndServe()
}
