package main

import (
	"flag"
	"net/http"
	"time"

	"github.com/golang/glog"
	"github.com/stianeikeland/go-rpio"
)

func main() {
	flag.Parse()

	rpio.Open()
	defer rpio.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		go openDoor()
		w.WriteHeader(http.StatusOK)
	})

	glog.Info("starting server")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		glog.Errorf("server shut down: %s", err)
	}
}

func openDoor() {
	pin := rpio.Pin(17)
	pin.Output()

	glog.Info("opening door")
	pin.High()
	time.Sleep(3 * time.Second)

	glog.Info("closing door")
	pin.Low()

	glog.Info("done")
}
