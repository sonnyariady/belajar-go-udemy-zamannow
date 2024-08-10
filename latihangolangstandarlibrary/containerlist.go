package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("Akang")
	data.PushBack("Sonny")
	data.PushFront("Profesor")

	var kepala *list.Element = data.Front()
	var buntut *list.Element = data.Back()

	fmt.Println("Kepala : ", kepala.Value)
	fmt.Println("Buntut: ", buntut.Value)

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
