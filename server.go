package main

import (
	"net/http"
	"flag"
	"math/rand"
	"time"
	"strconv"
)

type Service struct {
	Max int
}

func (s Service) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	ms := rand.Intn(s.Max) + 1
	time.Sleep(time.Duration(ms) * time.Millisecond)
	rw.Write([]byte(strconv.Itoa(ms)))
}

func main() {
	var max int
	flag.IntVar(&max, "max", 1, "The max sleep time in ms")

	flag.Parse()

	http.Handle("/", Service{Max: max})

	http.ListenAndServe(":80", nil)
}
