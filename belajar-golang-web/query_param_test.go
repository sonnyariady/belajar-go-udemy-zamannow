package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func SelamatSiang(writer http.ResponseWriter, request *http.Request) {
	nama := request.URL.Query().Get("nama")
	if nama == "" {
		fmt.Fprintln(writer, "Selamat siang")
	} else {
		fmt.Fprintf(writer, "Selamat siang %s \n", nama)
	}
}

func TestQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:2323/selamatSiang2?nama=Ahmad", nil)
	recorder := httptest.NewRecorder()

	SelamatSiang(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}

func MultipleParameter(writer http.ResponseWriter, request *http.Request) {
	nama := request.URL.Query().Get("nama")
	jurusan := request.URL.Query().Get("jurusan")
	fmt.Fprintf(writer, "Hallo mahasiswa baru bernama %s dari jurusan %s", nama, jurusan)

}

func TestQueryMultiParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:2323/pesantest?nama=Doni&jurusan=Manajemen", nil)
	recorder := httptest.NewRecorder()

	MultipleParameter(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}

func MultipleValueParameter(writer http.ResponseWriter, request *http.Request) {
	var query url.Values = request.URL.Query()
	nama := query.Get("nama")
	var foods []string = query["food"]
	fmt.Fprintf(writer, "%s sedang memakan %s", nama, strings.Join(foods, ","))

}

func TestQueryMultiValueParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:2323/pesantest?nama=Jaka&food=Siomay&food=Sate", nil)
	recorder := httptest.NewRecorder()

	MultipleValueParameter(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}
