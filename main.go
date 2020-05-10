package main

import (
	"fmt"
	"time"
)

func worker(queque []int, id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("loket", id, "started  queque", j)
		time.Sleep(time.Duration(queque[j-1]) * time.Second)
		fmt.Println("loket", id, "finished queque", j, queque[j-1], " second")
		results <- j
	}
}

func main() {
	// # Read integer
	var i int
	fmt.Print("Input jumlah loket: ")
	fmt.Scanf("%d", &i)

	//queque in second
	queque := []int{1, 2, 4, 2, 3, 5, 2, 3, 1, 3}
	numJobs := len(queque)
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= i; w++ {
		go worker(queque, w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}

}
