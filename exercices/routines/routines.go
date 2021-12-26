package routines

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup

func RoutinesWithAtomic() {
	var contador int64
	totaldegoroutines := 15

	var wg sync.WaitGroup
	wg.Add(totaldegoroutines)

	for i := 0; i < totaldegoroutines; i++ {
		go func() {
			atomic.AddInt64(&contador, 1)
			runtime.Gosched()
			fmt.Println("Contador:\t", atomic.LoadInt64(&contador))
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Valor final:", contador)
}

func RoutinesWithMutexAndRutime() {

	contador := 0
	totaldegoroutines := 150

	var wg sync.WaitGroup
	wg.Add(totaldegoroutines)

	var mu sync.Mutex

	for i := 0; i < totaldegoroutines; i++ {
		go func() {
			mu.Lock()
			v := contador
			runtime.Gosched()
			v++
			contador = v
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("Valor final:", contador)
}

func Routines(text string) {
	for i := 0; i < 10; i++ {
		fmt.Println(text, i)
		time.Sleep(200)
	}
	wg.Done()
}
