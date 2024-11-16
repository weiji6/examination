package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	M = 2000
	N = 10
	K = 200
)

func main() {
	station := make(chan int, K)
	totalPassengers := 0
	tripPassengers := make([]int, N)
	tripIndex := 0
	var wg sync.WaitGroup
	for i := 0; i < N; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			for {
				passenger, ok := <-station
				if !ok {
					fmt.Println(passenger)
					break
				}
				totalPassengers++
				tripPassengers[index]++
				if tripPassengers[index] == K || totalPassengers == M {
					fmt.Printf("Trip %d leaves with %d passengers\n", index+1, tripPassengers[index])
					tripPassengers[index] = 0
					tripIndex++
					if tripIndex >= N {
						tripIndex = 0
					}
					station <- 0
				}
			}
		}(i)
	}

	go func() {
		for i := 0; i < M; i++ {

			station <- 1

			time.Sleep(time.Millisecond * 10)
		}

		close(station)
	}()
	wg.Wait()
	fmt.Println("Final trip passenger counts:")
	for i, count := range tripPassengers {
		fmt.Printf("Trip %d: %d passengers\n", i+1, count)
	}
}
