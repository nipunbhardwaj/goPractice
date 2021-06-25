package main

import (
    "fmt"
    "log"
    "net/http"
    "io"
    "os"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	message, present := query["message"]
	if present && len(message) == 1 {
		fmt.Fprintf(w, "<h1>%s</h1>", message[0])
	} else {
		fmt.Fprintf(w, "<h1>Hello World !!!</h1>")
	}
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
	//get filename from path
	path := r.URL.Path

	// get current working directory
	wd, wderr := os.Getwd()
	if wderr != nil {
		log.Println(wderr)
		//File not found, send 404
		http.Error(w, "Something went wrong.", 404)
		return
	}
	done := make(chan bool)
	go func () {
		Openfile, err := os.Open( wd + "/views" + path)
		defer Openfile.Close()
		if err != nil {
			log.Println(err)
			//File not found, send 404
			http.Error(w, path + " no such file found.", 404)
			return
		}
		io.Copy(w, Openfile) //'Copy' the file to the client
		done <- true
	}()
	<-done
	return
}

func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	
	switch (path) {
	case "/":
		homeHandler(w,r)
	default:
		fileHandler(w,r)
	}
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":5000", nil))
}