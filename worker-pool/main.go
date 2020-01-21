package main

import (
	"fmt"
	"sync"
	"time"
)

type job struct {
	id     int
	number int
}

type result struct {
	job    job
	square int
}

var jobs = make(chan job, 10)
var results = make(chan result, 10)

func sq(x int) int {
	time.Sleep(2 * time.Second)
	return x * x
}

func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		out := result{job, sq(job.number)}
		results <- out
	}
	wg.Done()
}

func createworkerpool(noofworkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noofworkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

func allocate(noofworkers int) {
	for i := 0; i < noofworkers; i++ {
		out := job{i, i + 5}
		jobs <- out
	}
	close(jobs)
}

func readresult(done chan bool) {
	for result := range results {
		fmt.Printf("Job no: %d, input number: %d, result: %d\n", result.job.id, result.job.number, result.square)
	}

	done <- true
}

func main() {
	noofworkers := 10
	go allocate(noofworkers)

	done := make(chan bool)

	go readresult(done)
	createworkerpool(noofworkers)
	<-done
}
