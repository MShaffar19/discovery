package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/coreos/go-systemd/activation"
	"github.com/gorilla/mux"

	_ "github.com/coreos/discovery/http"
)

var addr = flag.String("addr", "", "web service address")

func main() {
	log.SetFlags(0)

	if addr != "" {
		http.ListenAndServe(*addr, nil)
	}

	listeners, err := activation.Listeners(true)
	if err != nil {
		panic(err)
	}

	if len(listeners) != 1 {
		panic("Unexpected number of socket activation fds")
	}

	http.Serve(listeners[0], nil)
}
