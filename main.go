package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("[*] Enter Targets To Scan(split them by ,): ")
	targetsInput, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln("Error reading input:", err)
	}
	targetsInput = strings.TrimSpace(targetsInput)

	fmt.Print("[*] Enter How Many Ports You Want To Scan: ")
	var numPorts int
	_, err = fmt.Scanf("%d", &numPorts)
	if err != nil {
		log.Fatalln("Error reading port number:", err)
	}

	targets := strings.Split(targetsInput, ",")

	if len(targets) > 1 {
		fmt.Println("[*] Scanning Multiple Targets")
		for _, target := range targets {
			target = strings.TrimSpace(target)
			if target != "" {
				fmt.Printf("\n[*] Scanning target: %s\n", target)
				scan(target, numPorts)
			}
		}
	} else {
		scan(targetsInput, numPorts)
	}
}

func scan(target string, numPorts int) {
	var wg sync.WaitGroup

	fmt.Printf("[*] Starting scan of %d ports\n", numPorts)

	semaphore := make(chan struct{}, 100)

	for port := 1; port <= numPorts; port++ {
		wg.Add(1)
		semaphore <- struct{}{}

		go func(p int) {
			defer func() {
				<-semaphore
				wg.Done()
			}()
			scanPort(target, p)
		}(port)
	}

	wg.Wait()
	fmt.Println("[*] Scan completed")
}

func scanPort(host string, port int) {
	target := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", target, 2*time.Second)

	if err != nil {
		if !strings.Contains(err.Error(), "refused") {
			fmt.Printf("Port %d error: %v\n", port, err)
		}
		return
	}

	defer conn.Close()
	fmt.Printf("[+] Port %d is open\n", port)
}
