package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
)

var count = 0

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func addHitCounter(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       // parse arguments, you have to call this by yourself
	fmt.Println(r.Form) // print form information in server side
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	// Save a copy of this request for debugging.
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, string(requestDump))
	fmt.Fprintf(w, r.RemoteAddr)
	count++

	time.Sleep(time.Second)

}

func cleanup() {
	//setup cleanup
	signal_chan := make(chan os.Signal, 1)
	signal.Notify(signal_chan, syscall.SIGINT)
	_ = <-signal_chan

	err := ioutil.WriteFile("count.txt", []byte(strconv.Itoa(count)), 0644)
	check(err)

	os.Exit(0)
}

func main() {
	go cleanup()

	http.HandleFunc("/add", addHitCounter) // set router
	http.HandleFunc("/favicon", nil)
	err := http.ListenAndServe(":8080", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
