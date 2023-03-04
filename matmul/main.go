package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const size = 250

var (
	A      = [size][size]int{}
	B      = [size][size]int{}
	C      = [size][size]int{}
	rwLock = sync.RWMutex{}
	cond   = sync.NewCond(rwLock.RLocker())
	wg     = sync.WaitGroup{}
)

func randMatrix(m *[size][size]int) {
	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			m[row][col] = rand.Intn(10) - 5
		}
	}
}

func calcRow(row int) {
	// row worker acquires read lock
	rwLock.RLock()
	for {
		// needed for this structure (multiple matrix mults in a loop)
		// sets up the worker
		// clears the wg.Wait at main() for loop
		wg.Done() 
		cond.Wait() // wait for signal from main after data is populated
		for col := 0; col < size; col++ {
			for i := 0; i < size; i++ {
				C[row][col] += A[row][i] * B[i][col]
			}
		}
	}
}

func main() {
	start := time.Now()
	// Add each row-wroker
	wg.Add(size)
	for row := 0; row < size; row++ {
		go calcRow(row) // span row workers
	}

	for k := 0; k < 100; k++ {
		wg.Wait() // wait for all row workers to be at cond.Wait()
		rwLock.Lock() // acquired write lock
		// write data into pre-allocated arrays
		randMatrix(&A)
		randMatrix(&B)
		// add workers to the wait group
		wg.Add(size)
		rwLock.Unlock() // unlock
		cond.Broadcast() // signal workers to work

	}
	elapsed := time.Since(start)
	fmt.Printf("multiplying 100 pairs random matrices too %s\n", elapsed)
}
