package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	data := ring.New(5)

	for i := 0; i < data.Len(); i++ {
		data.Value = "Value " + strconv.FormatInt(int64(i), 10)
		data = data.Next()
	}

	fmt.Println(data)

	data.Do(func(value interface{}) {
		fmt.Println(value)
	})

	var arrwarna = []string{"Merah", "Kuning", "Hijau", "Biru"}
	warnas := ring.New(4)

	for _, itm := range arrwarna {
		warnas.Value = itm
		warnas = warnas.Next()
	}
	fmt.Println(warnas)

	warnas.Do(func(value interface{}) {
		fmt.Println(value)
	})

	fmt.Println("Current val warnas 1:", warnas.Value)
	warnas = warnas.Next()
	fmt.Println("Current val warnas 2:", warnas.Value)
	warnas = warnas.Next()
	fmt.Println("Current val warnas 3:", warnas.Value)
	warnas = warnas.Next()
	fmt.Println("Current val warnas 4:", warnas.Value)
	warnas = warnas.Next()
	fmt.Println("Current val warnas 5:", warnas.Value)
}
