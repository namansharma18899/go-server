package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/about" {
		http.Error(w, "404 Path not found", http.StatusBadRequest)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not supported", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "hello from about!!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/about" {
		http.Error(w, "404 Path not found", http.StatusBadRequest)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not supported", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "hello from about!!")
}

func start_basic_server() {
	serverStarterConfig := make(map[string]string)
	serverStarterConfig["port"] = "8080"
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/form", formHandler)
	fmt.Printf("Starting server at port: ", serverStarterConfig["port"], "\n")
	server_address := fmt.Sprintf(":%v", serverStarterConfig["port"])
	if err := http.ListenAndServe(server_address, nil); err != nil {
		log.Fatal("Error while starting up the server!!")
	}
}
