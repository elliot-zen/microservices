package main

import (
	"errors"
	"log"
	"math/rand"
	"time"

	"github.com/sony/gobreaker"
)

func main() {
	cb := gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:        "demo",
		MaxRequests: 3, // Allowed number of requests for a half-open circuit;
		Timeout:     4, // Timeout for an open to haf-open transtion;
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			// This function decide on if the circuit will be open;
      log.Printf("=> ReadToTrip : failed(%d) total: (%d)", counts.TotalFailures, counts.Requests)
			failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
			return failureRatio >= 0.6
		},
		OnStateChange: func(name string, from, to gobreaker.State) {
			// This function is Executed on each state change;
			log.Printf("Circuit Breaker: %s, changed from %v to %v", name, from, to)
		},
	})

	for range 1000 {
		res, err := cb.Execute(func() (any, error) {
			res, isErr := randomProcess()
			if isErr {
				return nil, errors.New("error")
			}
			return res, nil
		})
    state := cb.State().String()
		if err != nil {
			log.Printf("[%s] Circuit breaker error %v", state, err)
		} else {
			log.Printf("[%s] Circuit breaker result %v", state, res)
		}
    time.Sleep(time.Second * 1)
	}
}

func randomProcess() (int, bool) {
	min := 10
	max := 30
	result := rand.Intn(max-min) + min
	return result, result < 25
}
