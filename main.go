package main

import (
	"flag"
	"log"
	"net"
	"sync"
	"time"
)

func worker(service string, timeout int) {
	timer := time.NewTimer(time.Duration(timeout) * time.Second)
	serviceAlive := false

	for !serviceAlive {
		_, err := net.Dial("tcp", service)
		
		if err != nil {
			log.Printf("Service %s not available yet\n", service)
			time.Sleep(1 * time.Second)
		} else {
			serviceAlive = true
			log.Printf("Service %s is available\n", service)
		}

		select {
		case <-timer.C:
			log.Fatalf("Service %s failed to respond within %d seconds\n", service, timeout)
		default:
			continue
		}
	}
}

func main() {
	timeoutOpt := flag.Int("timeout", 60, "Define how long should the program wait for each service to be available")
	
	flag.Parse()

	services := make([]string, 0)
	services = append(services, flag.Args()...)

	var wg sync.WaitGroup

	for _, service := range services {
		wg.Add(1)

		go func(service string) {
			defer wg.Done()
			worker(service, *timeoutOpt)
		}(service)
	}
	
	wg.Wait()
}