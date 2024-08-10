package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Dir("belajar/go/library/os.go"))
	fmt.Println(path.Dir("belajar/go/os.go"))

	fmt.Println(path.Base("belajar/go/library/math.go"))
	fmt.Println(path.Base("belajar/go/os.go"))

	fmt.Println(path.Ext("belajar/go/library/math.go"))
	fmt.Println(path.Ext("belajar/go/os.cs"))

	fmt.Println(path.Join("C:/Program Files", "Go", "sort.go"))
}
