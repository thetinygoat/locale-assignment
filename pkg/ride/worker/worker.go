package worker

import (
	"log"

	"github.com/adjust/redismq"
	"github.com/thetinygoat/localeai/pkg/entities/ride"
)

func worker(id int, jobs <-chan string, s ride.Service) {
	for j := range jobs {
		s.Dump(j)
	}
}

// StartWorkers spins up workers
func StartWorkers(n int, consumer *redismq.Consumer, s ride.Service) {
	bufferSize := 10
	jobs := make(chan string, bufferSize)
	defer close(jobs)
	for i := 1; i <= n; i++ {
		go worker(i, jobs, s)
	}
	for {
		m, err := consumer.Get()
		if err != nil {
			log.Fatal(err)
		}
		err = m.Ack()
		if err != nil {
			log.Fatal(err)
		}
		jobs <- m.Payload
	}

}
