package cmd

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/spf13/cobra"
	"golang.org/x/sync/semaphore"
)

var host string
var requestNumber int
var goroutineNumber int

func sendRequest(sem *semaphore.Weighted, ch chan<- time.Duration) {
	defer sem.Release(1)
	start := time.Now()
	_, err := http.Get(host)
	if err != nil {
		fmt.Println("Incorrect response", err)
	}
	ch <- time.Since(start)
}

func sumDurations(ch <-chan time.Duration, res chan<- time.Duration) {
	dur := time.Duration(0)
	for i := range ch {
		dur += i
	}

	res <- dur
}

var rootCmd = &cobra.Command{
	Use:   "ab",
	Short: "A benchmark",
	Long:  `A benchmark for a RESTful server.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		host = args[0]
		_, err := url.Parse(host)
		if err != nil {
			fmt.Println("Invalid host", err)
			return
		}

		ch := make(chan time.Duration, 1)
		res := make(chan time.Duration)
		go sumDurations(ch, res)

		sem := semaphore.NewWeighted(int64(goroutineNumber))
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		for i := 0; i < requestNumber; i++ {
			err = sem.Acquire(ctx, 1)
			if err != nil {
				fmt.Println("Cannot acquire semaphore", err)
				return
			}
			go sendRequest(sem, ch)
		}

		err = sem.Acquire(ctx, int64(goroutineNumber))
		if err != nil {
			fmt.Println("Cannot acquire semaphore", err)
			return
		}
		close(ch)

		dur := <-res
		close(res)
		fmt.Printf("Time to send %d requests with %d goroutines: %v\n", requestNumber, goroutineNumber, dur)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&requestNumber, "request_number", "n", 5, "Number of requests")
	rootCmd.PersistentFlags().IntVarP(&goroutineNumber, "goroutine_number", "c", 1, "Number of goroutines")
}
