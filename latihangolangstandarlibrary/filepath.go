package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.Dir("belajar/go/library/os.go"))
	fmt.Println(filepath.Base("belajar/go/library/os.go"))
	fmt.Println(filepath.Ext("belajar/go/library/os.go"))
	fmt.Println(filepath.IsLocal("belajar/go/library/os.go"))
	fmt.Println(filepath.IsLocal("c:\\belajar\\go\\library\\os.go"))
	fmt.Println(filepath.IsAbs("belajar/go/library/os.go"))
	fmt.Println(filepath.IsAbs("c:\\belajar\\go\\library\\os.go"))
	fmt.Println(filepath.Join("belajar", "go", "library", "os.go"))
}
