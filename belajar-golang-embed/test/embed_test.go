package test

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed appname.txt
var appname string

//go:embed kucing.jpg
var kucingbyte []byte

func TestString(t *testing.T) {

	fmt.Println(appname)
}

func TestByteArray(t *testing.T) {
	err := ioutil.WriteFile("kucing2.jpg", kucingbyte, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

//go:embed files/kucing.txt
//go:embed files/anjing.txt
//go:embed files/harimau.txt
//go:embed files/ayamjago.txt
var files embed.FS

func TestMultipleFiles(t *testing.T) {
	file1, _ := files.ReadFile("files/kucing.txt")
	fmt.Println(string(file1))

	file2, _ := files.ReadFile("files/anjing.txt")
	fmt.Println(string(file2))

	file3, _ := files.ReadFile("files/ayamjago.txt")
	fmt.Println(string(file3))

	file4, _ := files.ReadFile("files/harimau.txt")
	fmt.Println(string(file4))
}

//go:embed files/*.txt
var path embed.FS

func TestPathMatcher(t *testing.T) {
	dir, _ := path.ReadDir("files")
	for _, entry := range dir {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			content, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println(string(content))
		}
	}
}

//go:embed lirik/*.txt
var path2 embed.FS

func TestPathMatcher2(t *testing.T) {
	dir, _ := path2.ReadDir("lirik")
	for _, entry := range dir {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			content, _ := path2.ReadFile("lirik/" + entry.Name())
			fmt.Println(string(content))
		}
	}
}
