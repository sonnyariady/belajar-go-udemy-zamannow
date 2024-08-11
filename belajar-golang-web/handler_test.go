package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {

	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		//logic web
		fmt.Fprintln(writer, "Haai haaii ini halaman web saya")
	}

	server := http.Server{
		Addr:    "localhost:5678",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
