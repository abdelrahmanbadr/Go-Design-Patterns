package example2

import (
	"fmt"
	"time"
)

//We are explore the use of decorator pattern via following example in which we will extend an existing object that fetches a data from web service.
//We will decorate it by adding circuit breaker capabilities without changing the struct interface.
//Lets have a Fetcher interface that defines a contract for fetching some data from different sources.

// Args of fetching function
type Args map[string]string

// Data returned by fetch
type Data map[string]string

// Fetcher fetches a data from remote endpoint
type Fetcher interface {
	// Fetch fetches the data
	Fetch(args Args) (Data, error)
}

//A Retrier struct is the decorator that adds circuit breaker capabilities to any component that implements the Fetcher interface.
// The Retrier has a few properties that allow that. The RetryCount property defines the number of times that the retrier should try to fetch if there is an error. The WaitInterval property defines the interval between every retry. The Fetcher property is points to the object that is decorated.
// The Retrier calls the Fetch function of the decorated object until it succeeds or exceed the retry policy.
// Retrier retries multiple times
type Retrier struct {
	RetryCount   int
	WaitInterval time.Duration
	Fetcher      Fetcher
}

// Fetch fetches data
func (r *Retrier) Fetch(args Args) (Data, error) {
	for retry := 1; retry <= r.RetryCount; retry++ {
		fmt.Printf("Retrier retries to fetch for %d\n", retry)
		if data, err := r.Fetcher.Fetch(args); err == nil {
			fmt.Printf("Retrier fetched for %d\n", retry)
			return data, nil
		} else if retry == r.RetryCount {
			fmt.Printf("Retrier failed to fetch for %d times\n", retry)
			return Data{}, err
		}
		fmt.Printf("Retrier is waiting after error fetch for %v\n", r.WaitInterval)
		time.Sleep(r.WaitInterval)
	}

	return Data{}, nil
}

// A concrete implementation of the Fetcher interface is the Repository struct which provides some dummy data if the provided arguments are not empty, otherwise returns an error.
// The Repository struct is a concrete component in the context of The Decorator Pattern.
// Repository of data
type Repository struct{}

// Fetch fetches data
func (r *Repository) Fetch(args Args) (Data, error) {
	if len(args) == 0 {
		return Data{}, fmt.Errorf("No arguments are provided")
	}

	data := Data{
		"user":     "root",
		"password": "swordfish",
	}
	fmt.Printf("Repository fetched data successfully: %v\n", data)
	return data, nil
}

//Then we can add the new retry capabilities by wrapping the Repository instance with the Retrier
func main() {
	repository := &Repository{}
	retrier := &Retrier{
		RetryCount:   5,
		WaitInterval: time.Second,
		Fetcher:      repository,
	}

	data, err := repository.Fetch(Args{"id": "1"})
	fmt.Printf("#1 repository.Fetch: %v\n", data)

	data, err = retrier.Fetch(Args{})
	fmt.Printf("#2 retrier.Fetch error: %v\n", err)

	data, err = retrier.Fetch(Args{"id": "1"})
	fmt.Printf("#3 retrier.Fetch: %v\n", data)
}
