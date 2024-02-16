package main

import (
	"fmt"
	"os"
	"time"
	"os/signal"
	"syscall"
)

func producer(message chan string, done chan struct{}) {
	defer fmt.Println("Producer exiting...")
	for {
		select {
			case <-done:
				close(message)
				
				return
			default:
				select {
					case message <- "ping":
					default:
				}
		}
	}
}

func consumer(message chan string, done chan struct{}) {
	defer fmt.Println("Consumer exiting...")
	for {
		select {
			case <-done:
				return
			default:
				res, ok := <-message
				if(!ok) {
					return
				}else {
					fmt.Printf("Message is: %s\n", res);
				}
		}
	}
}

func main() {
	// Create a channel to receive signals and message channel
	sig := make(chan os.Signal, 1)
	message := make(chan string)
	done := make(chan struct{})

	// Notify the sig channel when a SIGINT (Ctrl+C) signal is received
	signal.Notify(sig, syscall.SIGINT)

	go producer(message,done)
	go consumer(message,done)

	// Block here until receive Ctrl+C singal
	<-sig

	// Signal received, handle the shutdown
	fmt.Println("\nCtrl+C received. Exiting...")

	defer func() {
		close(done)

		// Need waitgroups for remove this
		time.Sleep(time.Second)		
		fmt.Println("Done!")
	}()
}
