package concurrency

import (
	"fmt"
	"testing"
)

type Job int

func Test(t *testing.T) {
	jobs := make(chan Job)
	jobList := []Job{1, 2, 3}
	done := make(chan bool, len(jobList))

	go func() {
		for _, job := range jobList {
			jobs <- job
		}
		close(jobs)
	}()

	go func() {
		for job := range jobs {
			fmt.Println(job)
			done <- true
		}
	}()

	for i := 0; i < len(jobList); i++ {
		<-done
	}
}
