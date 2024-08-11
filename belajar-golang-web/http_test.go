package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Kucing(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Ngeong...ngeong")
}

func TestHttp(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:1212/kucing", nil)
	recorder := httptest.NewRecorder()

	Kucing(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}
