package main

import "flag"

var ip, port string

func init() {
	flag.StringVar(&ip, "ip", "127.0.0.1", "listen IP")
	flag.StringVar(&port, "port", "8080", "listen port")
	flag.Parse()
}

func main() {
	app := newApp(ip, port)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
