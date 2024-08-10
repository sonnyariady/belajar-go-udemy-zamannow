package belajar_golang_context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")
	contextG := context.WithValue(contextF, "g", "G")

	fmt.Println("context A: ", contextA)
	fmt.Println("context B: ", contextB)
	fmt.Println("context C: ", contextC)
	fmt.Println("context D: ", contextD)
	fmt.Println("context E: ", contextE)
	fmt.Println("context F: ", contextF)
	fmt.Println("context G: ", contextG)

	//get valuenya
	fmt.Println("Dapatin nilai c di konteks a ", contextA.Value("c")) //a tidak bisa ambil childnya
	fmt.Println("Dapatin nilai c di konteks g ", contextG.Value("c"))
	fmt.Println("Dapatin nilai f di konteks g ", contextG.Value("f"))
	fmt.Println("Dapatin nilai b di konteks g ", contextG.Value("b"))
	fmt.Println("Dapatin nilai f di konteks f ", contextF.Value("f"))
}

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second) //simulasi slow
			}

		}
	}()

	return destination
}

func TestContextWithCancel(t *testing.T) {
	fmt.Println("Jumlah goroutine awal: ", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := CreateCounter(ctx)
	fmt.Println("Jumlah goroutine current: ", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println("Counter: ", n)
		if n == 10 {
			break
		}
	}
	cancel() //kirim sinyal cancel ke context
	time.Sleep(2 * time.Second)

	fmt.Println("Jumlah goroutine akhir: ", runtime.NumGoroutine())
}

func TestContextWithTimeout(t *testing.T) {
	fmt.Println("Jumlah goroutine awal: ", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	destination := CreateCounter(ctx)
	fmt.Println("Jumlah goroutine current: ", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println("Counter: ", n)
	}

	time.Sleep(2 * time.Second)

	fmt.Println("Jumlah goroutine akhir: ", runtime.NumGoroutine())
}

func TestContextWithDeadline(t *testing.T) {
	fmt.Println("Jumlah goroutine awal: ", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(10*time.Second))
	defer cancel()

	destination := CreateCounter(ctx)
	fmt.Println("Jumlah goroutine current: ", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println("Counter: ", n)
	}

	time.Sleep(2 * time.Second)

	fmt.Println("Jumlah goroutine akhir: ", runtime.NumGoroutine())
}
