package main

import (
	"io"
	"log"
	"net/http"
	// must be enable for runtime memory profiling
	_ "net/http/pprof"
)

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func main() {
	http.HandleFunc("/hello", HelloServer)

	// if currently program not served as an http server,
	// we need to configure an http server for profiling purpose
	//go func() {
	//	http.ListenAndServe("0.0.0.0:6060", nil)
	//}()

	log.Fatal(http.ListenAndServe(":6060", nil))
}
