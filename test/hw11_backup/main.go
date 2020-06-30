package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

func main() {
	// do not rush to throw context down, think if it is useful with blocking operation?
	var timeoutString string
	flag.StringVar(&timeoutString, "timeout", "10s", "server connection timeout")
	flag.Parse()

	args := flag.Args()
	if len(args) < 2 {
		return
	}

	timeout, err := time.ParseDuration(timeoutString)
	if err != nil {
		log.Fatal(err)
	}

	address := net.JoinHostPort(args[0], args[1])
	client := NewTelnetClient(address, timeout, os.Stdin, os.Stdout)

	if err := client.Connect(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("...Connected to %s\n", address)
	defer func() {
		if err := client.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	var wg sync.WaitGroup
	wg.Add(2)

	// sig := make(chan os.Signal, 1)
	// signal.Notify(sig, syscall.SIGINT)

	done := make(chan struct{})
	go read(client, &wg, done)
	go write(client, &wg, done)

	wg.Wait()
}

func read(client TelnetClient, wg *sync.WaitGroup, done chan<- struct{}) {
	defer wg.Done()
	var err error
	for {
		err = client.Receive()
		if err != nil {
			fmt.Println("...Connection closed by remote server")
			close(done)
			break
		}
	}
}

func write(client TelnetClient, wg *sync.WaitGroup, done <-chan struct{}) {
	defer wg.Done()
	var err error
WRITE:
	for {
		select {
		case <-done:
			break WRITE
		default:
			err = client.Send()
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
