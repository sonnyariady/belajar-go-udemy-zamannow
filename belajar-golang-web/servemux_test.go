package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServeMux(t *testing.T) {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		//logic web
		fmt.Fprintln(writer, "Haai haaii ini halaman utama saya")
	})

	mux.HandleFunc("/kucing", func(writer http.ResponseWriter, request *http.Request) {
		//logic web
		fmt.Fprintln(writer, "Meong...meong")
	})

	mux.HandleFunc("/jawabarat/", func(writer http.ResponseWriter, request *http.Request) {
		//logic web
		fmt.Fprintln(writer, "Jawa Barat")
	})

	mux.HandleFunc("/jawabarat/bekasi/", func(writer http.ResponseWriter, request *http.Request) {
		//logic web
		fmt.Fprintln(writer, "Bekasi")
	})

	server := http.Server{
		Addr:    "localhost:1357",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestRequest(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		//logic web
		fmt.Fprintln(writer, request.Method)
		fmt.Fprintln(writer, request.RequestURI)
	}

	server := http.Server{
		Addr:    "localhost:2345",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
